package backend

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathSignTypedData(b *backend) *framework.Path {
	return &framework.Path{
		Pattern:      "accounts/" + framework.GenericNameRegex("name") + "/signTypeData",
		HelpSynopsis: "Sign a provided transaction object.",
		HelpDescription: `

    Sign a Type Data object with properties conforming to the Ethereum JSON-RPC documentation based on EIP-712.

    `,
		Fields: map[string]*framework.FieldSchema{
			"name": &framework.FieldSchema{Type: framework.TypeString},
			"typedData": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "(optional when creating new contract) The contract address the transaction is directed to.",
				Default:     "",
			},
		},
		ExistenceCheck: b.pathExistenceCheck,
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.CreateOperation: b.signTypedData,
		},
	}
}
