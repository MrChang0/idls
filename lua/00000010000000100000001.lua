local idls = require("IDLS")

function signalA(a,b)
-- write code here
-- don't delete this function
	idls.call("nodemcu","eventA","open")
end


function signalB(a,c)
-- write code here
-- don't delete this function
	idls.call("nodemcu","eventA","close")
end



    
    
    
   
    
    
    
    
    
    