$(document).ready(function () {

    $("#startGame").click(function(){
        var playerX = $( "#playerX" ).val();
        var playerO = $( "#playerO" ).val();
        $.post( "/xo/start",JSON.stringify({ playerX,playerO }), function( data,status,xhr) {
            if(xhr.status == 200){
                window.location="game.html"
            }
        });
    })
})