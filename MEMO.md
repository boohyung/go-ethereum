cmd/geth/main.go: geth
-> makeFullnode
-> startNode

cmd/geth/config.go:makeFullnode
-> makeConfigNode
-> RegisterEthService
-> shh, dashboard 

cmd/geth/config.go:MakeConfigNode
node/node.go: New
->makeAccountManager

cmd/geth/config.go:RegisterEthService
eth/backend.go: New 
-> CreateDB
-> NewBlloomIndexer
-> NewBlockChain
-> NewTxPool
-> NewProtocolManager
-> miner.New
-> gasprice.NewOracle

cmd/geth/main.go: geth->startNode ->Start
node/node.go: Start -> p2p.Server -> eth Start(shh,dash) ->startRPC 
eth/backend.go: Start -> protoclManager.Start

//protocol manager
// 프로토콜 매니져의 txs채널에서 이벤트가 오는지 감지하여 브로드캐스트 한다
eth/handler.go:Start->txBroadcastLoop(pm.txsCh) -> BroadcastTxs-> peer.SendTransactions
eth/peer.go: SendTransactions -> p2p.Send -> WriteMsg
p2p/message.go: WriteMsg -> ev.feed.Send
event/feed.go:ev.feed.Send -> channel.TrySend 


// 프로토콜 매니저의 이벤트 먹스에 마이닝 블록 이벤트를 구독한다
// 마이닝된 블록에 대한 구독채널에서 이벤트가 발생한경우
// 브로드캐스팅한다
eth/handler.go:Start->minedBroadcastLoop(core.NewMinedBlockSub) -> BroadcastBlock-> peer.SendTransactions -> peer.SendNewBlock
eth/peer.go: SendNewBlock-> p2p.send -> WriteMsg
p2p/message.go: WriteMsg -> ev.feed.Send
event/feed.go:ev.feed.Send -> channel.TrySend 


// 주기적으로 네트워크와 동기화 하고, 해시와 블록을 다운로드한다
// 해당블록을 모르는 피어에게 블록을 전달하고
 
eth/handler.go: Start->syncer
eth/sync.go: syncer -> pmfetcher.Start, syncronise 
eth/fetcher/fetcher.go: Start-> loop(queue pop) -> f.insert -> broadcastBloc k -> insertChain ->broadcastBlock 
core/blockchain.go: insertChain 
eth/handler.go: BroadcastBlock ->SendNewBlock, SendNewBlockHashes   
eth/peer.go: SendNewBlock -> p2p.Send
eth/peer.go: SendNewBlockHashes -> p2p.Send
eth/downloader/queue.go: queue


p2p/peer.go: handle(ping/pong) -> proto.in <-msg 
eth/handler.go: ReadMsg <-rw.in
eth/hanlder.go: handle -> handleMsg(NewBlockMsg)-> Enqueue -> f.inject
eth/fetcher/fetcher.go: loop -> f.inject -> enqueue -> queue push(block)
eth/hanlder.go: handle -> handleMsg(NewBlockHashesMsg)-> Enqueue -> Notify->f.notify
eth/fetcher/fetcher.go: loop -> f.notify -> rescheduleFetch 


// 이 함수는 새로운 커넥션에 대한 초기 트렌젝션을 관리한다
// 새로운 피어가 나타나면 현재까지 펜딩된 트렌젝션을 릴레이 한다
// 네트워크 밴드위스 관리를 위해 각 피어에 트렌젝션을 쪼개서 보낸다
eth/handler.go: Start->txsyncLoop
eth/sync.go: txsyncLoop -> pack.p.SendTransactions
eth/peer.go: SendTransactions -> p2p.Send -> WriteMsg
p2p/message.go: WriteMsg -> ev.feed.Send
event/feed.go:ev.feed.Send -> channel.TrySend 



accounts/abi/bind/base.go: DeployContract -> transact()-> transactor.SendTrasaction
internal/ethapi/api.go: SendTransacton->submitTransaction -> SendTx
eth/api_backend.go: SendTx -> AddLocal
core/tx_pool.go : AddLocal, AddRemove -> addTx -> pool.add



internal/ethapi/api.go: SendTransacton->submitTransaction
internal/ethapi/api.go: SendRawTransacton->submitTransaction




//트렌젝션이 공유되는 시퀀스
로컬(private account)이나 RPC(PublicTransactionPool)을 통해 SendTrasaction함수가 호출될 경우
local tx pool에 해당 트렌젝션을 추가하고 feed로 새로운 트렌젝션을 추가를 알리면,
노드 실행시 동작하던 broadcast 루프에서 감지하여 등록된 피어중 해당 트렌젝션을 모르는 피어에게
다시 SendTransaction을 하게 되고, 이는 p2p RLPx를 통해 인코딩 된후 TxMsg로 전송된다.
피어들은 생성시 동작시킨 handleMsg루프에서 해당 메시지를 읽고, pool.addRemote함수로 
자신의 풀에 해당 트렌젝션을 등록한다.


//노드간 통신연결
MakeFullNode에서 이더리움 프로토콜 매니져를 생성하면서
서브프로토콜로 P2P를 등록함(이때 p2p 핸들러로 Run함수를 저장)

StartNode에서 p2pServer가 스타트하게 되며 동시에 2가지 고루틴을 생성
1. Peer connection에 대한 Listening
2. dial(discovery:udp, transport: tcp)


새로운 노드의 dial이 성공하면 상대노드의 listning loop에서 accept를 하게 되고
양쪽 모두 setupConn 함수를 호출하게 되며
이때 peer를 생성하면서 p2p핸들러인 run함수가 호출된다

run함수 내부는
HandleMsg의 루프이며 ethereum 프로토콜 메시지가 전송되게 된다

// eth protocol message codes
const (
	// Protocol messages belonging to eth/62
	StatusMsg          = 0x00
	NewBlockHashesMsg  = 0x01
	TxMsg              = 0x02
	GetBlockHeadersMsg = 0x03
	BlockHeadersMsg    = 0x04
	GetBlockBodiesMsg  = 0x05
	BlockBodiesMsg     = 0x06
	NewBlockMsg        = 0x07

	// Protocol messages belonging to eth/63
	GetNodeDataMsg = 0x0d
	NodeDataMsg    = 0x0e
	GetReceiptsMsg = 0x0f
	ReceiptsMsg    = 0x10
)


//블록체인의 초기화/동기화
1. 노드 생성시
geth -> MakeFullNode -> regiseterEthService에서 이더리움 서비스 생성시
CreateDB로 체인DB를 초기화 한후, 메인넷의 제네시스 블록을 DB에 저장한다.
이상태로 새로운 퓨쳐블록을 처리하기 위한 5초짜리 티커를 설정하여 새로운 블록에 대한 처리를 시작한다
(이때 블록체인에는 블록이 하나) 

2. 피어와 연결되었을때/ 혹은 강제로(10초)
이후 이더리움 서비스 생성과정의 마지막 부분인 서브프로토콜생성과정에서 syncer라는 go 루틴이 동작하게 되는데
이 루틴에서 새로운 피어가 연결되었다는 메시지를 받거나/ 피어가 충분하지 않은 상황에서 10초 티커가 발생할 경우
상대 피어의 TD를 체크하여 나보다 크면 싱크를 시작한다
( 대부분의 경우 제네시스 블록에서 시작한 새로운 노드보다는 TD가 크다)
TD는 토탈 디피컬티로 해당 노드의 디피컬티가 크다는 것은 나보다 더 긴 캐노니컬 체인을 가지고 있다는 뜻으로 해석 가능하다.
downloader의 synchronise -> syncwithpeer -> spawnSync함수를 호춣하면
등록된 헤더패쳐, 바디패쳐, 영수증 페쳐, 헤더 프로세서가 동시에 실행되면서 
현재 내가 소유한 다음 블록부터 상대방의 체인을 읽어서 체인 DB에 저장하기 시작한다 



//마이닝의 시작과 블록이벤트 그리고 블록의 전파
eth의 StartNode 로부터 시작하여 이더리움 프로토콜이 시작하게 되면
크게 4개의 고루틴이 시작되는 데 그 중 하나가 MinedBroadcastLoop이다

geth -> StartNode -> Node.Start : cmd/geth/main Start -> p2p.Server -> eth Start (service.start) : node/node.go Start -> protoclManager.Start: eth/backend.go Start: eth/handler.go ->txBroadcastLoop(pm.txsCh) -> MinedBlockSub = pm.eventMux.Subscribe(core.NewMinedBlockEvent{}) -> minedBroadcastLoop(core.NewMinedBlockSub) ->syncer ->txsyncLoop

minedBroadcastLoop 함수안에서는
구독한 MinedBlockSub 채널에서 core.NewMinedBlockEvent가 발생하면
해당블록을 브로드캐스팅하도록 구성되어 있다

BroadcastBlock -> SendNewBlock(NewBlockMsg): eth/peer.go -> SendNewBlockHashes(NewBlockHashesMsg): eth/peer.go

core.NewMinedBlockEvent는 프로토콜 매니져 이후 실행되는 마이닝 루틴에서

StartNode-> StartMining: cmd/geth/main.go StartMining -> miner.Start : eth/backend.go minder.Start -> worker.start :miner/miner.go worker.start -> agent.start : miner/worker.go 여기서 agent는 ethash를 풀기위해 등록된 cpu agent임

워커가 블록을 발견했을 경우 노티파이 채널을 통해 먹스로 해당 이벤트를 포스팅하게된다

New -> miner.New : eth/backend.go New -> newWorker :miner/miner.go newWorker -> worker.wait -> self.mux.Post(core.NewMinedBlockEvent) : miner/worker.go

즉 마이닝을 실행한후 블록이 찾아지면 브로드캐스팅하는 것.
브로드 캐스팅을 수신한 피어의 프로토콜 매니져는
루프인 handleMsg에서 NewBlockMsg 수신하고 해당블록을 import하기위해
패쳐에 스케쥴을 enqueue한다 : eth/fetcher/fetcher.go

NewBlockHashesMsg 역시 같은 루프에서 처리하며
패쳐의 notify기능을 이용하여 announce된 블록리스트에 등록 한다

차후 패쳐의 메인루프에서 어나운스된 블록중 적절한 블록을 골라 체인에 삽입하고
매니저의 BroadcastBlock을 또다시 호출하게 됨으로서 처음 노드가 생성한 블록을 모르는
또다른 피어에게 해당 블록이 전달되게 된다

(A-B-C 형태로만 연결된 네트워크에서 A가 블록을 마이닝 한 경우 B에 전파되고,
B의 체인에 해당 블록이 삽입되면서 C로 브로드캐스팅되는 효과)




//마이닝 - 트렌젝션이 풀에서 블록에 포함되기 까지

eth.New->miner.New: eth/backend.go
New -> newWorker: miner/miner.go
newWorker: miner/worker.go

-> commitNewWork
이함수는 블록을 생성한다. 
먼저 새로운 헤더를 하나 생성한후 부모의 블록번호로 부터 해시정보를 얻어와 채우고
체인정보를 기반으로 난이도 필드를 초기화 한다

-> self.eth.TxPool().Pending()
-> types.NewTransactionsByPriceAndNonce(self.current.signer, pending) : core/types/transaction.go
트렌젝션풀에서는 트렌젝션을 2가지로 분류한다
pending: 현재 처리가능한 트렌젝션들
queue: 풀에 큐잉은 되었지만 처리가 불가능한 트렌젝션 (현제 체인 스테이트에서는 실행할수가 없어서)
이중 펜딩된 트렌젝션들을 가져와서 가격을 기준으로 정렬한다
이유는 블록에 포함시킬 트렌젝션들이 비싸야 채굴했을때 수익을 더 크게 가져갈수 있기 때문이다
(블록채굴 보상 + 트렌젝션에서 사용된 수수료)

-> commitTransactions -> commitTransaction 
-> ApplyTransaction: core/state_processor.go
트렌젝션을 하나하나 실행한다. 그리고 실행에 따른 스테이트를 DB에 업데이트 한다
또한 해당 트렌젝션에 대한 영수증을 만든다
트렌젝션이 실패할경우 검증을 시작하기 직전의 snapshot으로 돌려놓는다 
모든 트렌젝션의 실행이 완료될경우 mux.Post(core.PendingStateEvent{})를 호출한다

-> self.engine.Finalize: consensus/ethash/consensus.go
헤더의 루트를 완성하고 주어진 헤더 + 트렌젝션 + 엉클정보 + 영수증으로
블록을 만든다. 생성한 블록을 마이너 에이전트에게 전달한다

mine: miner.agent.go
에이전트는 블록이 담긴 워크를 마이닝하기 시작한다
self.engine.Seal: consensus/ethash/sealer.go 
ethash.mine
함수를 호출하게 되면 우리가 아는 pow기반으로 nonce를 찾기 시작한다.
논스를 찾았다면, 해당 결과를 worker 채널로 전송한다.

miner.wait에서 결과를 받으면
블록의 해시를 영수증에 업데이트 하고
self.chain.WriteBlockWithState 함수를 통해 체인에 해당 블록을 등록한다
그리고 나서 self.mux.Post(core.NewMinedBlockEvent{Block: block})
를 통해 블록을 브로드캐스팅한다 
self.chain.PostChainEvents(events, logs)

self.unconfirmed.Insert(block.NumberU64(), block.Hash())
해당블록을 unconfirmed 블록 리스트에 저장한다
unconfirmed란 아직 캐노니컬 체인에 등록되지 못한 블록들을 뜻한다
블록들은 unconfirmed상태에 존재하다가, 시간이지나면 miner.shift함수에 의해
제거된다 (이 schem은 유저가 지금까지 마이닝된 블록이 확률상
캐노니컬 채널에 등록될 확률이 높음을 알려줌으로서 
실제 채널에 등록되기전에 다른일을 할수있게 해준다)

















