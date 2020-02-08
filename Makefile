INSTALL := /usr/bin/install

vnet-platina-mk1:
	go build

install:
	$(INSTALL) -d $(DESTDIR)/usr/lib/goes
	$(INSTALL) vnet-platina-mk1 $(DESTDIR)/usr/lib/goes

.PHONY: vnet-platina-mk1 install
