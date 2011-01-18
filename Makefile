include $(GOROOT)/src/Make.inc
include $(GOROOT)/src/Make.cmd

CLEANFILES=goping *.$O

ICMP_SRC=src/icmp/icmp.go src/icmp/ping.go
ICMP_PKG=icmp
ICMP_BIN=$(ICMP_PKG).$O

MAIN_SRC=src/main.go
MAIN_PKG=main
MAIN_BIN=$(MAIN_PKG).$O

BIN_NAME=goping

all: $(BIN_NAME)

$(BIN_NAME): $(MAIN_BIN)
	$(QUOTED_GOBIN)/$(LD) -L . -o $(BIN_NAME) $(MAIN_BIN)

$(MAIN_BIN): $(ICMP_BIN) $(MAIN_SRC)
	$(QUOTED_GOBIN)/$(GC) -I . -o $(MAIN_BIN) $(MAIN_SRC)

$(ICMP_BIN): $(ICMP_SRC)
	$(QUOTED_GOBIN)/$(GC) -o $(ICMP_BIN) $(ICMP_SRC)

format:
	gofmt -w=true $(ICMP_SRC) $(MAIN_SRC)
