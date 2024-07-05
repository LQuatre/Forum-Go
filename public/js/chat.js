var $messages = $(".messages-content"),
  d,
  h,
  m,
  i = 0;

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (data) {
  thismessage = JSON.parse(data.data);
  console.log("message: " + JSON.stringify(thismessage));
  if (thismessage.typing) {
    wsLoadingMessage(thismessage);
  } else if (thismessage.message) {
    wsMessage(thismessage);
  } else {
    wsLoadingMessage(thismessage);
  }
};

$(window).on("load", function () {
  $messages.mCustomScrollbar();
  ws.send(
    JSON.stringify({
      username: $("#username").val(),
      typing: false,
    })
  );
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

  const message = {
    username: $("#username").val(),
    message: msg,
  };

  ws.send(JSON.stringify(message));
}

$(".message-submit").click(function () {
  insertMessage();
});

var typing = false;

$(window).on("keydown", function (e) {
  if (!typing) {
    typing = true;
    message = {
      username: $("#username").val(),
      typing: true,
    };
    ws.send(JSON.stringify(message));
    setTimeout(function () {
      typing = false;
      message = {
        username: $("#username").val(),
        typing: false,
      };
      ws.send(JSON.stringify(message));
    }, 1000);
  }
  if (e.which == 13) {
    typing = false;
    insertMessage();
    return false;
  }
});

function wsLoadingMessage(data) {
  if (data.username == $("#username").val()) {
    return;
  }
  if (data.typing == true) {
    $(
      '<div class="message loading new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure><span></span></div>'
    ).appendTo($(".mCSB_container"));
    updateScrollbar();
  } else if (data.typing == false) {
    $(".message.loading").remove();
  }
}

function wsMessage(data) {
  if (data.username == $("#username").val()) {
    return;
  }
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

$(".chatBtn").click(function () {
  $(".chat").css("display", "flex");
  $(".chatBtn").css("display", "none");
});

$(".chat-close").click(function () {
  $(".chat").css("display", "none");
  $(".chatBtn").css("display", "flex");
});
