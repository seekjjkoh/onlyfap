return require('luvit')(function (...)
  local http = require('http')

  http.createServer(function (req, res)
    local body = "Hello world\n"
    res:setHeader("Content-Type", "text/plain")
    res:setHeader("Content-Length", #body)
    res:finish(body)
  end):listen(1337, '0.0.0.0')
  
  print('Server running at port: 1337/')
end, ...)
