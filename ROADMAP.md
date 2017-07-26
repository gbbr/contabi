ROADMAP:
- Set up front-end from scratch (app/ui) starting with routes (https://github.com/reactjs/redux/tree/master/examples/real-world) pointing to basic empty pages
- [DONE] Find a way to serve & build BE & FE together w/ live-reload (webpack + go-bindatassetfs, etc) from the same binary. Perhaps with github.com/sqs/rego
- Set up Go dummy backend with dummy state for each route
- Set up auth (JWT?)
- Learn GraphQL and create a server (which app will call into for queries and will later query a gRPC server)
- Explore application until data structures start to take shape

TODO:
- Find a way to commit files inside /vendor
- Find a way to run go-bindata in -dev using a flag or env var (currently -dev is hardcoded in ui package)
