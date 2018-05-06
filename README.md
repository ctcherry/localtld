# localtld

Tiny DNS server for local development

## Usage

The program is configured via environment variables.

* `TLD` - The TLD that the DNS server will respond to. Queries that fall outside of this TLD will fail.
* `IP` - The IP in the A record that all responses will have.
* `LISTEN` - The IP and port binding that the DNS server will listen on.

### Usage Examples

    ./localtld

Runs a DNS server on 127.0.0.1 port 10053 that responds to all DNS requests with an A record of 127.0.0.1

    TLD=foo ./localtld

Runs a DNS server on 127.0.0.1 port 10053 that responds to all DNS requests under the TLD foo, with an A record of 127.0.0.1

    IP=10.0.0.1 ./localtld

Runs a DNS server on 127.0.0.1 port 10053 that responds to all DNS requests, with an A record of 10.0.0.1

    LISTEN=0.0.0.0:11153 ./localtld

Runs a DNS server on 0.0.0.0 port 11153 that responds to all DNS requests, with an A record of 127.0.0.1
