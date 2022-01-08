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
				LockingList: []types.Locking{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LockingCount: 2,
				InspectionList: []types.Inspection{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				InspectionCount: 2,
				TransferenceList: []types.Transference{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TransferenceCount: 2,
				HouseList: []types.House{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				HouseCount: 2,
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
			desc: "duplicated certificate",
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
			valid: false,
		},
		{
			desc: "invalid certificate count",
			genState: &types.GenesisState{
				CertificateList: []types.Certificate{
					{
						Id: 1,
					},
				},
				CertificateCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated locking",
			genState: &types.GenesisState{
				LockingList: []types.Locking{
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
			desc: "invalid locking count",
			genState: &types.GenesisState{
				LockingList: []types.Locking{
					{
						Id: 1,
					},
				},
				LockingCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated inspection",
			genState: &types.GenesisState{
				InspectionList: []types.Inspection{
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
			desc: "invalid inspection count",
			genState: &types.GenesisState{
				InspectionList: []types.Inspection{
					{
						Id: 1,
					},
				},
				InspectionCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated transference",
			genState: &types.GenesisState{
				TransferenceList: []types.Transference{
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
			desc: "invalid transference count",
			genState: &types.GenesisState{
				TransferenceList: []types.Transference{
					{
						Id: 1,
					},
				},
				TransferenceCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated house",
			genState: &types.GenesisState{
				HouseList: []types.House{
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
			desc: "invalid house count",
			genState: &types.GenesisState{
				HouseList: []types.House{
					{
						Id: 1,
					},
				},
				HouseCount: 0,
			},
			valid: false,
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
