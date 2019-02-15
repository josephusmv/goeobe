function BindMouseEvents(){
    $(document).on("dragstart", function() {
     return false;
    });
    //Drag drop is triggered from Mouse Down as the first event.
    $(".leaf-note").mousedown(function(e){
        resetAllZIndex();
        $(document).off("mousedown");
        var posXS = 0, posYS = 0;   //start position
        var posXM = 0, posYM = 0;   //moved distance
        
        //use hover style when move
        $(".leaft-notes-link").addClass("leaft-notes-link-move");
        
        //elmnt = $(this).parent();
        elmnt = $(this);
        e = e || window.event;
        
        // get the mouse cursor position at startup:
        posXS = e.clientX;
        posYS = e.clientY;
        console.log("Mouse move over from (" + e.clientX + ", " + e.clientY + ")");
        
        $(document).mouseup(function(e){
            /* stop moving when mouse button is released:*/
            $(document).off("mouseup");
            $(document).off("mousemove");
            $(document).on("mousedown");
            elmnt.css("z-index", "5");
            console.log("Mouse move over To (" +newPosX + ", " + newPosY + ")");
            $(".leaft-notes-link").removeClass("leaft-notes-link-move");
        });
        
        // call a function whenever the cursor moves:
        $(document).mousemove(function(e){
            elmnt.css("z-index", "5");
            e = e || window.event;
            
            console.log("---> Element from Pos (" +e.clientX + ", " + e.clientY + ")");
            //calculate the move distance
            posXM = e.clientX - posXS;
            posYM = e.clientY - posYS;
            //save new start pos
            posXS = e.clientX;
            posYS = e.clientY;
            
            // set the element's new position:
            curPos = elmnt.offset()
            newPosX = curPos.left + posXM;
            newPosY = curPos.top + posYM;
            
            //calculate container frame border
            fOff = $(".frameDIV").offset();
            if(newPosY < fOff.top + 8)
                newPosY=fOff.top + 8;
            if(newPosX < fOff.left + 8)
                newPosX=fOff.left + 8;
             
            eHeight = elmnt.height();
            fHeight = $(".frameDIV").height();
            if(newPosY + eHeight > fOff.top + fHeight)
                newPosY=fOff.top + fHeight - eHeight;
            eWidth = elmnt.width();
            fWidth = $(".frameDIV").width();
            if(newPosX + eWidth > fOff.left + fWidth)
                newPosX=fOff.top + fWidth - eWidth;
            
            
            elmnt.offset({top: newPosY, left: newPosX});
            console.log("---> Try to Move Element to new Pos (" +newPosX + ", " + newPosY + ")");
        });
    }); // $(".titlearea").mousedown(function(e){
    
    function resetAllZIndex(){
        $(".frameDIV").children(".leaf-note").each(function(){
            $(this).css("z-index", "0")
        });
    }
        
    /*
    //editarea auto expand
    $(".datacontent").keydown(function(){
        console.log("datacontent keydown");
        var el = $(this)[0];
        setTimeout(function(){
            if(el.scrollHeight >= 230 ){
                el.style.cssText = 'height:auto; padding:0';
                el.style.cssText = 'height:12.8em';
                el.style.overflow = "auto";
                return;
            }
            
            el.style.cssText = 'height:auto; padding:0';
            el.style.cssText = 'height:' + el.scrollHeight + 'px';
            el.style.overflow = "hidden";
        },0);        
    });
    */
        
    //titlearea edit
    $(".leaf-note").dblclick(function(){
        alert("Enter Detain page...");
    });
    
}

//should triggered by onload and sidebar buttons.
function calculateFrameDivSize(sidebarcollapsed){
    var winHeight = $(window).height();
    var navHeight = $("#navbarHolder").height() + 10; //add 10 pixels for margin and padding    
    var winWidth = $(window).width();
    var sidbWidth = $("#sidebar").width() + 10;//add 10 pixels for margin and padding 
    
    var expctwidth = winWidth;
    if (sidebarcollapsed){
        expctwidth = winWidth - sidbWidth;
    }
    
    
    $(".frameDIV").height(winHeight-navHeight);
    $(".frameDIV").width(expctwidth);
}