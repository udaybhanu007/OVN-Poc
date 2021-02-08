package config

type Configuration struct {
	DbConfig DBConfig `json:"dbConfig"`
}

type DBConfig struct {
	CassandraHost        string `json:"cassandraHost"`
	CassandraKeyspace    string `json:"cassandraKeyspace"`
	CassandraConsistency string `json:"cassandraConsistency"`
}
