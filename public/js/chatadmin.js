var $messages = $(".messages-content"),
  d,
  h,
  m,
  i = 0;

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (data) {
  message = JSON.parse(data.data);
  console.log(message);
  if (message.typing) {
    wsTyping(message);
  } else if (message.message) {
    wsMessage(JSON.parse(data.data));
  } else {
    $(".message.loading").remove();
  }
};

$(window).on("load", function () {
  $messages.mCustomScrollbar();
});

function updateScrollbar() {
  $messages.mCustomScrollbar("update").mCustomScrollbar("scrollTo", "bottom", {
    scrollInertia: 10,
    timeout: 0,
  });
}

function setDate() {
  d = new Date();
  if (m != d.getMinutes()) {
    m = d.getMinutes();
    $('<div class="timestamp">' + d.getHours() + ":" + m + "</div>").appendTo(
      $(".message:last")
    );
  }
}

function insertMessage() {
  msg = $(".message-input").val();
  if ($.trim(msg) == "") {
    return false;
  }
  $('<div class="message message-personal">' + msg + "</div>")
    .appendTo($(".mCSB_container"))
    .addClass("new");
  setDate();
  $(".message-input").val(null);
  updateScrollbar();

  
  message = {
    username: "system",
    user_uuid: $("#useruuid").val(),
    message: msg,
    typing: false,
    destination: $("#destination").val()
  };

  ws.send(JSON.stringify(message));
}

$(".message-submit").click(function () {
  insertMessage();
});

var typing = false;
var hetypedagain = false;
var typingTimer = null;

$(window).on("keydown", function (e) {
  if (!typing) {
    typing = true;
    if (hetypedagain) {
      return;
    }
    ws.send(JSON.stringify({
      username: "system",
      user_uuid: $("#useruuid").val(),
      typing: true,
      destination: $("#destination").val()
    }));
  }
  hetypedagain = true;

  if (typing) {
    clearTimeout(typingTimer);
    typingTimer = setTimeout(function () {
      typing = false;
      hetypedagain = false;
      ws.send(JSON.stringify({
        username: "system",
        user_uuid: $("#useruuid").val(),
        destination: $("#destination").val(),
        typing: false
      }));
    }, 2000);
  }

  if (e.which == 13) {
    typing = false;
    insertMessage();
    return false;
  }
});

function wsTyping(data) {
  if (data.username == $("#username").val() && data.destination == $("#destination").val()) {
    return;
  }
  if (data.destination == $("#useruuid").val()) {
    $(
      '<div class="message loading new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure><span></span></div>'
    ).appendTo($(".mCSB_container"));
    updateScrollbar();
  }
}

function wsMessage(data) {
  if (data.username == $("#username").val() && data.destination == $("#destination").val()) {
    return;
  }
  if (data.destination == $("#useruuid").val()) {
    $(".message.loading").remove();
    $(
      '<div class="message new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure>' +
        data.message +
        "</div>"
    )
      .appendTo($(".mCSB_container"))
      .addClass("new");
    setDate();
    updateScrollbar();
  }
}