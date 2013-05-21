
/*
 * GET home page.
 */

exports.index = function(req, res){
   res.sendfile("views/index.html",function(err){ // Transfer The File With COntent Type Text/HTML
        if(err){
            res.send(err);
        }else{
            res.end(); // Send The Response
        }
    })
};