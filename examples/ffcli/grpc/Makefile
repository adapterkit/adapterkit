.PHONY: adapterkit
adapterkit:
	adapterkit generate -mod github.com/pmg-tools/adapterkit-module-swissknife -tpl templates/ proto/swissknife.proto
	adapterkit generate -mod github.com/Doozers/adapterkit-module-tictac -tpl templates/ proto/tictac.proto


.PHONY: clean
clean:
	rm -rf pkg/*_gen
