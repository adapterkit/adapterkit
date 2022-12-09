.PHONY: adapterkit
adapterkit:
	adapterkit generate -mod github.com/pmg-tools/adapterkit-module-swissknife -tpl ../../../templates/ffcli/unix-cmd proto/swissknife.proto
	adapterkit generate -mod github.com/Doozers/adapterkit-module-tictac -tpl ../../../templates/ffcli/unix-cmd proto/tictac.proto

.PHONY: clean
clean:
	rm -rf pkg/*_gen
