This is the "Basecoin" example application built on the Cosmos-SDK.  This
"Basecoin" is not affiliated with [Coinbase](http://www.getbasecoin.com/), nor
the [stable coin](http://www.getbasecoin.com/).

You need a recent version of `glide` to install Basecoin's dependencies.

Install `glide` by running

```bash
> go get github.com/Masterminds/glide
```

or change to `cosmos-sdk` directory and run

```bash
> make get_tools
```

Then, you can build the cmd binaries (NOTE: a work in progress!), or run the tests.

```bash
> make get_vendor_deps
> make build
> make test
```

If you want to create a new application, start by copying the Basecoin app.
