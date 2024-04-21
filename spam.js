var pesan = prompt("masukan pesan di sini","");
var serangan = parseInt(prompt("mau berapa kali bos?",10));

window.InputEvent = window.Event || window.InputEvent;
var event = new InputEvent("input",{bubbles:true});

var text = document.getElementsByClassName("_13NKt copyable-text selectable-text")[1];

for (let i=0;i<serangan;i++) {
    text.innerHTML = pesan;
    text.dispatchEvent(event);
    document.getElementsByClassName("_4sWnG")[0].click()
}