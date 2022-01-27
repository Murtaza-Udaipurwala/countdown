PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin

all: countdown

countdown:
	go build

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 countdown $(DESTDIR)$(BINDIR)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/countdown

clean:
	rm -f countdown

.PHONY: all install uninstall clean
