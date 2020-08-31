package pkg

const ClientSlave = 1 << 0

type ReplicationMain struct {
	Host string
	Port int
}

type ReplicationReplica struct {
	Host string
	Port int
}

func NewReplicationMain() *ReplicationMain {
	return new(ReplicationMain)
}
func replicationFeedSlaves() {

}

func (m *ReplicationMain) ReplicationUnsetMaster() {

}
