.PHONY: adapterkit
adapterkit:
	protoc -I. --gotemplate_out=template_dir=templates,debug=true:. proto/swissknife.proto


.PHONY: clean
clean:
	rm -rf pkg/*_gen
