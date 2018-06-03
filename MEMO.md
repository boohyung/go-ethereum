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
