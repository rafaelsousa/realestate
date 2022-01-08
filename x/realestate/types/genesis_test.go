package types_test

import (
	"testing"

	"github.com/rafaelsousa/realestate/x/realestate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				PropertyList: []types.Property{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PropertyCount: 2,
				CertificateList: []types.Certificate{
	{
		Id: 0,
	},
	{
		Id: 1,
	},
},
CertificateCount: 2,
// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated property",
			genState: &types.GenesisState{
				PropertyList: []types.Property{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid property count",
			genState: &types.GenesisState{
				PropertyList: []types.Property{
					{
						Id: 1,
					},
				},
				PropertyCount: 0,
			},
			valid: false,
		},
		{
	desc:     "duplicated certificate",
	genState: &types.GenesisState{
		CertificateList: []types.Certificate{
			{
				Id: 0,
			},
			{
				Id: 0,
			},
		},
	},
	valid:    false,
},
{
	desc:     "invalid certificate count",
	genState: &types.GenesisState{
		CertificateList: []types.Certificate{
			{
				Id: 1,
			},
		},
		CertificateCount: 0,
	},
	valid:    false,
},
// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
