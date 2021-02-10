EXE=app

start:
	go build -o $(EXE)
	./$(EXE)
