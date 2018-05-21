// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// geth is the official command-line client for Ethereum.
package main //go 언어에서 main 패키지로 선언하면 실행파일이 생성된다고 한다

import (
	//go 시스템 패키지, 코드에서 사용시 어떤 용도로 쓰이는지 업데이트 하자.
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	//이더리움 내부 패키지
	//다른 폴더의 소스코드를 패키지화해서 사용한다
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/console"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/debug"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/node"

	//gopkg.in은 stable api라는걸 지원하는데, 정확하진 않지만 github에서 버전별로 패키지를 받아 사용할수 있는것 같다
	//urfave/cli 패키지는 command line 기반의 go application을 만들기 위해 사용되는 패키지.
	//지원하는 기능은 콘솔창, 입력된 커맨드 파싱 및 콜백함수, 설정파일 읽어 변수초기화 하기 등등
	"gopkg.in/urfave/cli.v1"
)

const (
	//go에서 상수를 선언할때 괄호로 묶어서 사용한 케이스
	//타입이 명시적이지 않으면 자동으로 타입 추론
	//const clientIdentifier string = "geth" 랑 같음
	clientIdentifier = "geth" // Client identifier to advertise over the network
)

var (
	//go에서 변수를 묶어서 선언함

	// Git SHA1 commit hash of the release (set via linker flags)
	gitCommit = ""
	// The app that holds all commands and flags.
	//utils패키지의 NewApp함수를 사용한다. 해당함수는 utils/flags.go에 선언되어 있음. 리턴값은 cli.App* 이고 내부적으로 cli.NewApp을 호출함
	//cli패키지는 urfave 주석의 내용대로 커맨드라인 app을 만들기 위한 helper
	app = utils.NewApp(gitCommit, "the go-ethereum command line interface")
	// flags that configure the node
	// cli의 Flag type형태 배열을 선언함 
	// 내부적으로는 string, bool, directory 내용들이 섞여 있음
	// util의 flags.go에 실제 cli의 flag type선언 참조
	nodeFlags = []cli.Flag{
		utils.IdentityFlag,
		utils.UnlockedAccountFlag,
		utils.PasswordFileFlag,
		utils.BootnodesFlag,
		utils.BootnodesV4Flag,
		utils.BootnodesV5Flag,
		utils.DataDirFlag,
		utils.KeyStoreDirFlag,
		utils.NoUSBFlag,
		utils.DashboardEnabledFlag,
		utils.DashboardAddrFlag,
		utils.DashboardPortFlag,
		utils.DashboardRefreshFlag,
		utils.EthashCacheDirFlag,
		utils.EthashCachesInMemoryFlag,
		utils.EthashCachesOnDiskFlag,
		utils.EthashDatasetDirFlag,
		utils.EthashDatasetsInMemoryFlag,
		utils.EthashDatasetsOnDiskFlag,
		utils.TxPoolNoLocalsFlag,
		utils.TxPoolJournalFlag,
		utils.TxPoolRejournalFlag,
		utils.TxPoolPriceLimitFlag,
		utils.TxPoolPriceBumpFlag,
		utils.TxPoolAccountSlotsFlag,
		utils.TxPoolGlobalSlotsFlag,
		utils.TxPoolAccountQueueFlag,
		utils.TxPoolGlobalQueueFlag,
		utils.TxPoolLifetimeFlag,
		utils.FastSyncFlag,
		utils.LightModeFlag,
		utils.SyncModeFlag,
		utils.GCModeFlag,
		utils.LightServFlag,
		utils.LightPeersFlag,
		utils.LightKDFFlag,
		utils.CacheFlag,
		utils.CacheDatabaseFlag,
		utils.CacheGCFlag,
		utils.TrieCacheGenFlag,
		utils.ListenPortFlag,
		utils.MaxPeersFlag,
		utils.MaxPendingPeersFlag,
		utils.EtherbaseFlag,
		utils.GasPriceFlag,
		utils.MinerThreadsFlag,
		utils.MiningEnabledFlag,
		utils.TargetGasLimitFlag,
		utils.NATFlag,
		utils.NoDiscoverFlag,
		utils.DiscoveryV5Flag,
		utils.NetrestrictFlag,
		utils.NodeKeyFileFlag,
		utils.NodeKeyHexFlag,
		utils.DeveloperFlag,
		utils.DeveloperPeriodFlag,
		utils.TestnetFlag,
		utils.RinkebyFlag,
		utils.VMEnableDebugFlag,
		utils.NetworkIdFlag,
		utils.RPCCORSDomainFlag,
		utils.RPCVirtualHostsFlag,
		utils.EthStatsURLFlag,
		utils.MetricsEnabledFlag,
		utils.FakePoWFlag,
		utils.NoCompactionFlag,
		utils.GpoBlocksFlag,
		utils.GpoPercentileFlag,
		utils.ExtraDataFlag,
		configFileFlag,
	}

	rpcFlags = []cli.Flag{
		utils.RPCEnabledFlag,
		utils.RPCListenAddrFlag,
		utils.RPCPortFlag,
		utils.RPCApiFlag,
		utils.WSEnabledFlag,
		utils.WSListenAddrFlag,
		utils.WSPortFlag,
		utils.WSApiFlag,
		utils.WSAllowedOriginsFlag,
		utils.IPCDisabledFlag,
		utils.IPCPathFlag,
	}

	whisperFlags = []cli.Flag{
		utils.WhisperEnabledFlag,
		utils.WhisperMaxMessageSizeFlag,
		utils.WhisperMinPOWFlag,
	}
)

func init() {
	// Initialize the CLI app and start Geth

	app.Action = geth //app이 생성될때 불리우는 main함수.
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2013-2018 The go-ethereum Authors"
	app.Commands = []cli.Command{ //app에서 지원할 command들
	//TODO:각 커맨드에 대한 내용 추가
		// See chaincmd.go:
		initCommand,
		importCommand,
		exportCommand,
		importPreimagesCommand,
		exportPreimagesCommand,
		copydbCommand,
		removedbCommand,
		dumpCommand,
		// See monitorcmd.go:
		monitorCommand,
		// See accountcmd.go:
		accountCommand,
		walletCommand,
		// See consolecmd.go:
		consoleCommand,
		attachCommand,
		javascriptCommand,
		// See misccmd.go:
		makecacheCommand,
		makedagCommand,
		versionCommand,
		bugCommand,
		licenseCommand,
		// See config.go
		dumpConfigCommand,
	}
	//TODO:정렬을 왜하지?
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = append(app.Flags, nodeFlags...)
	app.Flags = append(app.Flags, rpcFlags...)
	app.Flags = append(app.Flags, consoleFlags...)
	app.Flags = append(app.Flags, debug.Flags...)
	app.Flags = append(app.Flags, whisperFlags...)

	//Before는 sub command가 실행되기 전 불리우는 함수
	//모든 명령어를 실행하기전 카운터와 가스 리밋값을 설정한다.
	app.Before = func(ctx *cli.Context) error {
		//go의 멀티cpu지원 - not multi-thread
		runtime.GOMAXPROCS(runtime.NumCPU())
		if err := debug.Setup(ctx); err != nil {
			return err
		}
		// Start system runtime metrics collection
		//3초간 동작중인 프로세스의 다양한 메트릭을 수집하여
		//Counter나 Gauge형태로 사용하는것 같다
		//TODO:어떤 용도로 쓰는지는 차후확인
		//metics/metrics.go 파일을 참조
		go metrics.CollectProcessMetrics(3 * time.Second)

		//네트워크를 셋업한다
		//TargetGasLimit값을 설정함
		//utils/flags.go
		utils.SetupNetwork(ctx)
		return nil
	}
	//Before는 sub command가 실행된 후 불리우는 함수

	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}
}

func main() {
	//cli app을 실행한다. 당연히 action에 해당하는 geth함수가 불릴것.
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// geth is the main entry point into the system if no special subcommand is ran.
// It creates a default node based on the command line arguments and runs it in
// blocking mode, waiting for it to be shut down.
func geth(ctx *cli.Context) error {
	//config.go에 정의됨
	//makeFullNode 에선 아래와 같은 순서로 노드를 초기화함
	//makeConfigNode
		//노드생성 - 노드는 서비스들이 등록될수 있는 컨테이너
		//node config - p2p, ipc, http, WebSocket(ws), Node Userident
		//eth config - 이더베이스(account), gas price(gpo), txpool, ethash
	//RegisterEthService - 노드스택에 full 혹은 light ethereum 서비스를 등록
	//registerDashboardSurf - 이더리움의 data visualizer인 대시보드 서비스를 붙인다 
	//enableWhisper - 최소 1개의 whisper flag가 있다면 enable 
	//registerEthStatsService -URL이 있을경우 이더 stat daemon에 등록한다
	node := makeFullNode(ctx)
	//노드 시작
	startNode(ctx, node)
	node.Wait()
	return nil
}

// startNode boots up the system node and all registered protocols, after which
// it unlocks any requested accounts, and starts the RPC/IPC interfaces and the
// miner.
func startNode(ctx *cli.Context, stack *node.Node) {
	debug.Memsize.Add("node", stack)

	// Start up the node itself
	// node/node.go
	// 노드가 실행되면 p2p서버를 초기화 하고 
	// 노드에 등록했던 서비스프로토콜을 실행한다
	// RPC실행(admin, debug, web3)
	// 인터럽트를 기다려 node를 종료시키는 스레드 동시실행
	utils.StartNode(stack)

	// Unlock any account specifically requested
	// keystore를 참조하여 account를 언락한다
	//TODO: 키스토어 관련 좀더 분석할것
	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	passwords := utils.MakePasswordList(ctx)
	unlocks := strings.Split(ctx.GlobalString(utils.UnlockedAccountFlag.Name), ",")
	for i, account := range unlocks {
		if trimmed := strings.TrimSpace(account); trimmed != "" {
			unlockAccount(ctx, ks, trimmed, i, passwords)
		}
	}
	// Register wallet event handlers to open and auto-derive wallets
	events := make(chan accounts.WalletEvent, 16)
	stack.AccountManager().Subscribe(events)

	go func() {
		// Create a chain state reader for self-derivation
		// 체인정보를 읽는 rpcClient
		rpcClient, err := stack.Attach()
		if err != nil {
			utils.Fatalf("Failed to attach to self: %v", err)
		}
		stateReader := ethclient.NewClient(rpcClient)

		// Open any wallets already attached
		// 등록된 지갑들을 open한다.
		for _, wallet := range stack.AccountManager().Wallets() {
			if err := wallet.Open(""); err != nil {
				log.Warn("Failed to open wallet", "url", wallet.URL(), "err", err)
			}
		}
		// Listen for wallet event till termination
		for event := range events {
			switch event.Kind {
			case accounts.WalletArrived:
				if err := event.Wallet.Open(""); err != nil {
					log.Warn("New wallet appeared, failed to open", "url", event.Wallet.URL(), "err", err)
				}
			case accounts.WalletOpened:
				status, _ := event.Wallet.Status()
				log.Info("New wallet appeared", "url", event.Wallet.URL(), "status", status)

				if event.Wallet.URL().Scheme == "ledger" {
					event.Wallet.SelfDerive(accounts.DefaultLedgerBaseDerivationPath, stateReader)
				} else {
					event.Wallet.SelfDerive(accounts.DefaultBaseDerivationPath, stateReader)
				}

			case accounts.WalletDropped:
				log.Info("Old wallet dropped", "url", event.Wallet.URL())
				event.Wallet.Close()
			}
		}
	}()
	// Start auxiliary services if enabled
	if ctx.GlobalBool(utils.MiningEnabledFlag.Name) || ctx.GlobalBool(utils.DeveloperFlag.Name) {
		// Mining only makes sense if a full Ethereum node is running
		if ctx.GlobalBool(utils.LightModeFlag.Name) || ctx.GlobalString(utils.SyncModeFlag.Name) == "light" {
			utils.Fatalf("Light clients do not support mining")
		}
		var ethereum *eth.Ethereum
		if err := stack.Service(&ethereum); err != nil {
			utils.Fatalf("Ethereum service not running: %v", err)
		}
		// Use a reduced number of threads if requested
		if threads := ctx.GlobalInt(utils.MinerThreadsFlag.Name); threads > 0 {
			type threaded interface {
				SetThreads(threads int)
			}
			if th, ok := ethereum.Engine().(threaded); ok {
				th.SetThreads(threads)
			}
		}
		// Set the gas price to the limits from the CLI and start mining
		ethereum.TxPool().SetGasPrice(utils.GlobalBig(ctx, utils.GasPriceFlag.Name))
		//mining시작
		if err := ethereum.StartMining(true); err != nil {
			utils.Fatalf("Failed to start mining: %v", err)
		}
	}
}
