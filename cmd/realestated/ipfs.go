package main

import (
	"context"
	"fmt"
	config "github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader" // This package is needed so that all the preloaded plugins are loaded automatically
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SetupIPFS() {
	fmt.Println("##### Starting IPFS node ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Spawn a node using the default path (~/.ipfs), assuming that a repo exists there already
	fmt.Println("#### Spawning node on default repo ")
	_, err := spawnDefault(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawnDefault node: %s", err))
	}

	fmt.Println("### IPFS node is running")
}

// setupPlugins initializes the IPFS Repo
func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

/// ------ Spawning the node
// Creates an IPFS node and returns its coreAPI
func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		switch err.(type) {
		case fsrepo.NoRepoError:
			err := createNewRepo()
			if err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}
	// Construct the node
	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

func createNewRepo() error {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	repoDir := fmt.Sprintf("%s/.ipfs", homeDir)
	err = os.Mkdir(repoDir, 0700)
	if err != nil {
		return err
	}
	cfg, err := config.Init(ioutil.Discard, 2048)
	err = fsrepo.Init(repoDir, cfg)
	if err != nil {
		return err
	}
	return nil
}

// Spawns a node on the default repo location, if the repo exists
func spawnDefault(ctx context.Context) (icore.CoreAPI, error) {
	defaultPath, err := config.PathRoot()
	if err != nil {
		return nil, err
	}

	if err := setupPlugins(defaultPath); err != nil {
		return nil, err

	}

	return createNode(ctx, defaultPath)
}
