package cmd

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"
	hidnodecli "github.com/hypersign-protocol/hid-node/x/ssi/client/cli"
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
)

func extendDebug(debugCmd *cobra.Command) *cobra.Command {
	debugCmd.AddCommand(
		ed25519Cmd(),
		secp256k1Cmd(),
		bbsCmd(),
		bjjCmd(),
		signSSIDocCmd(),
	)
	return debugCmd
}

func bjjCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bjj",
		Short: "BabyJubJub debug commonds",
	}

	cmd.AddCommand(bjjRandomCmd())

	return cmd
}

func bbsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bbs",
		Short: "bbs debug commands",
	}

	cmd.AddCommand(
		blsRandomCmd(),
	)

	return cmd
}

func secp256k1Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secp256k1",
		Short: "secp256k1 debug commands",
	}

	cmd.AddCommand(
		secp256k1RandomCmd(),
		secp256k1Bech32AddressCmd(),
		secp256k1EthRandomCmd(),
	)

	return cmd
}

func blsRandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random blsBbs keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, privKey, err := bbs.GenerateKeyPair(sha256.New, nil)
			if err != nil {
				return err
			}

			// Convert Public Key Object to Multibase
			pubKeyBytes, err := pubKey.Marshal()
			if err != nil {
				return err
			}

			publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
			if err != nil {
				return err
			}

			// Convert Private Object to Bytes
			privKeyBytes, err := privKey.Marshal()
			if err != nil {
				return err
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_base_64"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_base_64"`
			}{
				PubKeyBase64:    base64.StdEncoding.EncodeToString(pubKeyBytes),
				PubKeyMultibase: publicKeyMultibase,
				PrivKeyBase64:   base64.StdEncoding.EncodeToString(privKeyBytes),
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

func bjjRandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random BabyJubJub keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Generate Key Pair
			babyJubjubPrivKey := babyjub.NewRandPrivKey()
			babyJubjubPubKey := babyJubjubPrivKey.Public()

			// Prepare Priv key
			var privKeyBytes [32]byte = babyJubjubPrivKey
			var privKeyHex string = hex.EncodeToString(privKeyBytes[:])

			// Prepare Public Key
			var pubKeyHex string = babyJubjubPubKey.Compress().String()

			// Prepare Multibase Public Key
			pubKeyBytes, err := hex.DecodeString(pubKeyHex)
			if err != nil {
				panic(err)
			}
			pubKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
			if err != nil {
				panic(err)
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_hex"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_hex"`
			}{
				PubKeyBase64:    pubKeyHex,
				PubKeyMultibase: pubKeyMultibase,
				PrivKeyBase64:   privKeyHex,
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

func secp256k1RandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random secp256k1 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKeyObj := secp256k1.GenPrivKey()

			privateKey := privateKeyObj.Bytes()
			publicKeyCompressed := privateKeyObj.PubKey().Bytes()

			publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKeyCompressed)
			if err != nil {
				panic(err)
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_base_64"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_base_64"`
			}{
				PubKeyBase64:    base64.StdEncoding.EncodeToString(publicKeyCompressed),
				PubKeyMultibase: publicKeyMultibase,
				PrivKeyBase64:   base64.StdEncoding.EncodeToString(privateKey),
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

func secp256k1Bech32AddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bech32-addr [base64-encoded-public-key] [prefix]",
		Short: "Converts a compressed base64 encoded secp256k1 public key to bech32 address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			base64PubKey := args[0]
			addrPrefix := args[1]

			publicKeyBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
			if err != nil {
				panic(err)
			}

			bech32address := publicKeyToBech32Address(addrPrefix, publicKeyBytes)

			_, err = fmt.Fprintln(cmd.OutOrStdout(), bech32address)
			return err
		},
	}

	return cmd
}

func secp256k1EthRandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eth-hex-random",
		Short: "Generate random Ethereum hex-encoded secp256k1 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKeyObj := secp256k1.GenPrivKey()
			privateKey := privateKeyObj.Bytes()

			publicKeyCompressed := privateKeyObj.PubKey().Bytes()

			publicKeyUncompressed, err := ethercrypto.DecompressPubkey(publicKeyCompressed)
			if err != nil {
				return err
			}
			ethereumAddress := ethercrypto.PubkeyToAddress(*publicKeyUncompressed).Hex()

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_hex"`
				PrivKeyBase64   string `json:"priv_key_hex"`
				EthereumAddress string `json:"ethereum_address"`
			}{
				PubKeyBase64:    hex.EncodeToString(publicKeyCompressed),
				PrivKeyBase64:   hex.EncodeToString(privateKey),
				EthereumAddress: ethereumAddress,
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

// ed25519Cmd returns cobra Command.
func ed25519Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ed25519",
		Short: "ed25519 debug commands",
	}

	cmd.AddCommand(
		ed25519RandomCmd(),
		base64toMultibase58Cmd(),
		multibase58toBase64Cmd(),
	)

	return cmd
}

// Generate signature for Schema and VC Status Document
func signSSIDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-ssi-doc",
		Short: "Get signature for the signed document",
	}

	cmd.AddCommand(
		signDidDocCmd(),
		signSchemaDocCmd(),
		signCredStatusDocCmd(),
	)
	return cmd
}

func signDidDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "did-doc [doc] [private-key] [proof-object-without-signature]",
		Short: "Did Document signature",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argDidDoc := args[0]
			argPrivateKey := args[1]
			argProofObjectWithoutSignature := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal DID Document
			var didDoc types.DidDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDoc), &didDoc)
			if err != nil {
				return err
			}

			// Unmarshal Proof Object
			var didDocProof types.DocumentProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argProofObjectWithoutSignature), &didDocProof)
			if err != nil {
				return err
			}

			// Sign DID Document
			var signature string
			switch didDocProof.Type {
			case types.Ed25519Signature2020:
				var didDocBytes []byte
				if len(didDoc.Context) > 0 {
					didDocBytes, err = ldcontext.Ed25519Signature2020Normalize(&didDoc, &didDocProof)
					if err != nil {
						return err
					}
				} else {
					didDocBytes = didDoc.GetSignBytes()
				}

				signature, err = hidnodecli.GetEd25519Signature2020(argPrivateKey, didDocBytes[:])
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1Signature2019:
				var didDocBytes []byte
				if len(didDoc.Context) > 0 {
					didDocBytes, err = ldcontext.EcdsaSecp256k1Signature2019Normalize(&didDoc, &didDocProof)
					if err != nil {
						return err
					}
				} else {
					didDocBytes = didDoc.GetSignBytes()
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1Signature2019(argPrivateKey, didDocBytes[:])
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1RecoverySignature2020:
				var didDocBytes []byte
				if len(didDoc.Context) > 0 {
					didDocBytes, err = ldcontext.EcdsaSecp256k1RecoverySignature2020Normalize(&didDoc, &didDocProof)
					if err != nil {
						return err
					}
				} else {
					didDocBytes = didDoc.GetSignBytes()
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1RecoverySignature2020(argPrivateKey, didDocBytes[:])
				if err != nil {
					return err
				}
			case types.BbsBlsSignature2020:
				var didDocBytes []byte
				if len(didDoc.Context) > 0 {
					didDocBytes, err = ldcontext.BbsBlsSignature2020Normalize(&didDoc, &didDocProof)
					if err != nil {
						return err
					}
				} else {
					didDocBytes = didDoc.GetSignBytes()
				}

				signature, err = hidnodecli.GetBbsBlsSignature2020(argPrivateKey, didDocBytes[:])
				if err != nil {
					return err
				}
			case types.BabyJubJubSignature2023:
				signature, err = hidnodecli.GetBabyJubJubSignature2023(argPrivateKey, didDoc.GetSignBytes())
				if err != nil {
					return err
				}
			default:
				panic("recieved unsupported signing-algo. Supported algorithms are: [Ed25519Signature2020, EcdsaSecp256k1Signature2019, EcdsaSecp256k1RecoverySignature2020, BbsBlsSignature2020, BabyJubJubSignature2023]")
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

func signSchemaDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema-doc [doc] [private-key] [proof-object-without-signature]",
		Short: "Schema Document signature",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argSchemaDoc := args[0]
			argPrivateKey := args[1]
			argProofObjectWithoutSignature := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Schema Document
			var schemaDoc types.CredentialSchemaDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaDoc), &schemaDoc)
			if err != nil {
				return err
			}
			
			// Unmarshal Proof Object
			var credSchemaDocProof types.DocumentProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argProofObjectWithoutSignature), &credSchemaDocProof)
			if err != nil {
				return err
			}

			// Sign Schema Document
			var signature string
			switch credSchemaDocProof.Type {
			case types.Ed25519Signature2020:
				credSchemaDocBytes, err := ldcontext.Ed25519Signature2020Normalize(&schemaDoc, &credSchemaDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEd25519Signature2020(argPrivateKey, credSchemaDocBytes)
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1Signature2019:
				credSchemaDocBytes, err := ldcontext.EcdsaSecp256k1Signature2019Normalize(&schemaDoc, &credSchemaDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1Signature2019(argPrivateKey, credSchemaDocBytes)
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1RecoverySignature2020:
				credSchemaDocBytes, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Normalize(&schemaDoc, &credSchemaDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1RecoverySignature2020(argPrivateKey, credSchemaDocBytes)
				if err != nil {
					return err
				}
			case types.BbsBlsSignature2020:
				credSchemaDocBytes, err := ldcontext.BbsBlsSignature2020Normalize(&schemaDoc, &credSchemaDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetBbsBlsSignature2020(argPrivateKey, credSchemaDocBytes)
				if err != nil {
					return err
				}
			case types.BabyJubJubSignature2023:
				signature, err = hidnodecli.GetBabyJubJubSignature2023(argPrivateKey, schemaDoc.GetSignBytes())
				if err != nil {
					return err
				}
			default:
				panic(fmt.Sprintf(
					"recieved unsupported signing-algo '%v'. Supported algorithms are: [Ed25519Signature2020, EcdsaSecp256k1Signature2019, EcdsaSecp256k1RecoverySignature2020, BbsBlsSignature2020, BabyJubJubSignature2023]",
					credSchemaDocProof.Type,
				))
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

func signCredStatusDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cred-status-doc [doc] [private-key] [proof-object-without-signature]",
		Short: "Credential Status Document signature",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCredStatusDoc := args[0]
			argPrivateKey := args[1]
			argProofObjectWithoutSignature := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Credential Status Document
			var credStatusDoc types.CredentialStatusDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argCredStatusDoc), &credStatusDoc)
			if err != nil {
				return err
			}

			// Unmarshal Proof Object
			var credStatusDocProof types.DocumentProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argProofObjectWithoutSignature), &credStatusDocProof)
			if err != nil {
				return err
			}

			// Sign Credential Status Document
			var signature string
			switch credStatusDocProof.Type {
			case types.Ed25519Signature2020:
				credStatusDocBytes, err := ldcontext.Ed25519Signature2020Normalize(&credStatusDoc, &credStatusDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEd25519Signature2020(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1Signature2019:
				credStatusDocBytes, err := ldcontext.EcdsaSecp256k1Signature2019Normalize(&credStatusDoc, &credStatusDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1Signature2019(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case types.EcdsaSecp256k1RecoverySignature2020:
				credStatusDocBytes, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Normalize(&credStatusDoc, &credStatusDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetEcdsaSecp256k1RecoverySignature2020(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case types.BbsBlsSignature2020:
				credStatusDocBytes, err := ldcontext.BbsBlsSignature2020Normalize(&credStatusDoc, &credStatusDocProof)
				if err != nil {
					return err
				}

				signature, err = hidnodecli.GetBbsBlsSignature2020(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case types.BabyJubJubSignature2023:
				signature, err = hidnodecli.GetBabyJubJubSignature2023(argPrivateKey, credStatusDoc.GetSignBytes())
				if err != nil {
					return err
				}
			default:
				panic("recieved unsupported signing-algo. Supported algorithms are: [Ed25519Signature2020, EcdsaSecp256k1Signature2019, EcdsaSecp256k1RecoverySignature2020, BbsBlsSignature2020, BabyJubJubSignature2023]")
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

// ed25519Cmd returns cobra Command.
func ed25519RandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random ed25519 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
			if err != nil {
				return err
			}

			// A 2-byte header must be prefixed before Ed25519 public key based on
			// W3C's Ed25519VerificationKey2020 Specification for the attribute `publicKeyMultibase`
			// Read more: https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/#ed25519verificationkey2020
			var publicKeyWithHeader []byte
			publicKeyWithHeader = append(publicKeyWithHeader, append([]byte{0xed, 0x01}, pubKey...)...)

			keyInfo := struct {
				PubKeyMultibase  string `json:"pub_key_multibase"`
				PrivKeyMultibase string `json:"priv_key_base_64"`
			}{
				PubKeyMultibase: "z" + base58.Encode(publicKeyWithHeader),
				// W3C's Ed25519VerificationKey2020 Specification has not explicitly mentioned about the encoding of the private key
				// or whether it should be prefixed similar to the public key. For now, the encoding of private key remains Base64
				// with no prefix
				PrivKeyMultibase: base64.StdEncoding.EncodeToString(privKey),
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

// Converts base-64 encoded ed25519 public key to multibase58 string
func base64toMultibase58Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base64-multibase58 [base64-encoded-public-key]",
		Short: "Convert base64 string to multibase58 string",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			base64Str := args[0]
			bytes, err := base64.StdEncoding.DecodeString(base64Str)
			if err != nil {
				return err
			}

			multibase58Str, err := multibase.Encode(multibase.Base58BTC, bytes)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), multibase58Str)
			return err
		},
	}

	return cmd
}

// Converts multibase58 encoded ed25519 public key to base-64 string
func multibase58toBase64Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multibase58-base64 [multibase58-public-key]",
		Short: "Convert multibase58 string to base64 string",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			multibase58Str := args[0][1:]
			multibase58Bytes := base58.Decode(multibase58Str)

			base64Str := base64.StdEncoding.EncodeToString(multibase58Bytes)

			_, err := fmt.Fprintln(cmd.OutOrStdout(), base64Str)
			return err
		},
	}

	return cmd
}
