var table = document.getElementById("calendar");
    function tdclick(event){
      var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", '/forCheck', false ); // false for synchronous request
    xmlHttp.send( null );
    console.log(xmlHttp.responseText);
    // event.stopPropagation()
};
var htmlstring = "";
for (var i = 1; i <= 5; i++) {
//open tr tag
htmlstring += "<tr>";
for (var i = 1; i <= 30; i++) {
  var days = "<td onclick='tdclick(event);'>" + i + "</td>";
  //print 1 to 31 with td tag
  htmlstring += days;
  //if i divide by 7 and remainder is 0
  if (i % 7 == 0 || i == 31) {
    htmlstring += "</tr>";
    htmlstring += "<tr>";
  }
};
htmlstring += "</tr>";
};
//console.log(htmlstring);
table.insertAdjacentHTML('beforeend', htmlstring)