GO=go
RM=rm -rf
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
SOURCE=github.com/omidplus/arvan
OUTPUT=bin
CLIENT=client
SERVER=server

all: client server

client:
	$(GOBUILD) -o $(OUTPUT)/$(CLIENT) $(SOURCE)/$(CLIENT)

server:
	$(GOBUILD) -o $(OUTPUT)/$(SERVER) $(SOURCE)/$(SERVER)

clean:
	${GOCLEAN}
	$(RM) $(OUTPUT)/$(CLIENT)
	$(RM) $(OUTPUT)/$(SERVER)
	$(RM) $(OUTPUT)

.PHONY: client server clean

