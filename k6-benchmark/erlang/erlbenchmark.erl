-module(erlbenchmark).
-export([
    main/1,
    start/1
]).

main(_) ->
    start(2345).

start(Port) ->
    {ok, Sock} = gen_tcp:listen(Port, [{active, false}]), 
    io:format("Server listening to 2345~n"),
    accept(Sock).

accept(Listen) ->
    {ok, Socket} = gen_tcp:accept(Listen),
    handle(Socket),
    accept(Listen).

handle(Conn) ->
    gen_tcp:send(Conn, response("Hello World")),
    gen_tcp:close(Conn).

response(Str) ->
    B = iolist_to_binary(Str),
    iolist_to_binary(
      io_lib:fwrite(
         "HTTP/1.0 200 OK\nContent-Type: text/html\nContent-Length: ~p\n\n~s",
         [size(B), B])).