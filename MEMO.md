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

