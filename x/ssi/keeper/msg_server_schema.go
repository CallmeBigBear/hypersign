package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

func storeCredentialSchema(
	k msgServer,
	goCtx context.Context,
	schemaDoc *types.CredentialSchemaDocument,
	schemaProof *types.DocumentProof,
	txAuthor string,
	msgTypeUrl string,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schemaID := schemaDoc.GetId()
	chainNamespace := k.GetChainNamespace(&ctx)

	// Validate Document Proof
	if err := schemaProof.Validate(); err != nil {
		return err
	}

	// Get the Did Document of Schema's Author and check if Author's DID is deactivated
	authorDidDocumentState, err := k.getDidDocumentState(&ctx, schemaDoc.GetAuthor())
	if err != nil {
		return sdkerrors.Wrap(err, fmt.Sprintf("unable to get author`s DID %s from store", schemaDoc.GetAuthor()))
	}
	if authorDidDocumentState.DidDocumentMetadata.Deactivated {
		return sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to create schema", authorDidDocumentState.DidDocument.Id))
	}

	// Check if Schema ID is valid
	if err := verification.IsValidID(schemaID, chainNamespace, "schemaDocument"); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidSchemaID, err.Error())
	}

	// Check if Schema already exists
	if k.hasCredentialSchema(ctx, schemaID) {
		return sdkerrors.Wrap(types.ErrSchemaExists, fmt.Sprintf("Schema ID:  %s", schemaID))
	}

	// Check if the `name` attribute of schema is in PascalCase
	if !isStringInPascalCase(schemaDoc.Name) {
		return sdkerrors.Wrapf(types.ErrInvalidCredentialSchema, "name must always be in PascalCase: %v", schemaDoc.Name)
	}

	// Check if `properties` field is a valid JSON document
	if !isValidJSONDocument(schemaDoc.Schema.Properties) {
		return sdkerrors.Wrapf(types.ErrInvalidCredentialSchema, "properties must be a valid JSON document: %v", schemaDoc.Schema.Properties)
	}

	// Signature check
	if err := k.VerifyDocumentProof(ctx, schemaDoc, schemaProof); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidClientSpecType, err.Error())
	}

	var schema = types.CredentialSchemaState{
		CredentialSchemaDocument: schemaDoc,
		CredentialSchemaProof:    schemaProof,
	}

	k.setCredentialSchemaInStore(ctx, schema)

	// Emit a successful Schema Registration event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(msgTypeUrl, sdk.NewAttribute("tx_author", txAuthor)),
	)

	return nil
}

func (k msgServer) RegisterCredentialSchema(goCtx context.Context, msg *types.MsgRegisterCredentialSchema) (*types.MsgRegisterCredentialSchemaResponse, error) {
	if err := storeCredentialSchema(k, goCtx, msg.CredentialSchemaDocument, msg.CredentialSchemaProof, msg.TxAuthor, msg.Type()); err != nil {
		return nil, err
	}

	return &types.MsgRegisterCredentialSchemaResponse{}, nil
}

func (k msgServer) UpdateCredentialSchema(goCtx context.Context, msg *types.MsgUpdateCredentialSchema) (*types.MsgUpdateCredentialSchemaResponse, error) {
	if err := storeCredentialSchema(k, goCtx, msg.CredentialSchemaDocument, msg.CredentialSchemaProof, msg.TxAuthor, msg.Type()); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCredentialSchemaResponse{}, nil
}

// isStringInPascalCase checks if an input string is in Pascal case or not
func isStringInPascalCase(s string) bool {
	pascalCaseRegex := regexp.MustCompile(`^[A-Z][a-zA-Z0-9]*$`)

	return pascalCaseRegex.MatchString(s)
}

// isValidJSONDocument checks if the input string is a valid JSON string
func isValidJSONDocument(s string) bool {
	var jsonStruct map[string]interface{}
	return json.Unmarshal([]byte(s), &jsonStruct) == nil
}
