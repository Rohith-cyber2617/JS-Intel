package config

type Options struct {
	URL         string
	List        string
	Output      string
	Threads     int
	Depth       int
	Endpoints   bool
	FoundOnly   bool
	Verify      bool
	RandomAgent bool
	Silent      bool
	Update      bool
	Help        bool
}
