$(document).ready(function () {
    $.getJSON("/xo/info", function (data) {
        $("#playerX").text(data.playerX.name) 
        $("#playerO").text(data.playerO.name)
        $("#turnOf").text(data.turnOf.name) 
        for(row = 0;row<data.grid.length;row++){
            for(column = 0;column<data.grid[row];column++){
                $(`#grid-${row}-${column}`).text(grid[row][column])
            }
        }
    })

})