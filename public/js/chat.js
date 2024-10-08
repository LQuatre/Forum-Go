var $messages = $(".messages-content"),
  d,
  h,
  m,
  i = 0;

var responses = {}

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
  setTimeout(function() {
    fakeMessage();
  }, 100);
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

  // add the message to response [0]
  if (responses[0] === undefined) {
    console.log("response is undefined 0");
    responses[0] = msg;
    
    setTimeout(function() {
      console.log(responses[0]);
      fakeMessage();
    }, 1000 + (Math.random() * 20) * 100);
    return;
  }

  if (responses[1] === undefined) {
    console.log("response is undefined 1");
    responses[1] = msg;
    
    setTimeout(function() {
      console.log(responses[1]);
      fakeMessage();
    }, 1000 + (Math.random() * 20) * 100);
    return;
  }

  if (responses[2] === undefined) {
    console.log("response is undefined 2 ");
    responses[2] = msg;

    // send the response to the server with POST url
    $.ajax({
      url: "/chatbot/createticket",
      type: "POST",
      data: {
        name: responses[1],
        description: responses[2],
        user_uuid: $("#useruuid").val(),
      },
      success: function(data) {
        console.log(data);
      }
    });

    setTimeout(function() {
      console.log(responses[2]);
      fakeMessage();
    }, 1000 + (Math.random() * 20) * 100);
    return;
  }

  
  // message = {
  //   username: $("#username").val(),
  //   user_uuid: $("#useruuid").val(),
  //   message: msg,
  //   typing: false,
  //   destination: $("#destination").val()
  // };

  // ws.send(JSON.stringify(message));
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
      username: $("#username").val(),
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
        username: $("#username").val(),
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
  console.log(data.destination);
  console.log($("#useruuid").val());
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

var Fake = [
  'Bonjour, je suis la pour vous aider ! Dites moi quel est votre problème ?',
  'Quel seras le titre de votre ticket ?',
  'Pouvez-vous me donner une description de votre problème ?',
  'Votre demande à bien été prise en compte. Nous allons vous aider.',
]

function fakeMessage() {
  if ($('.message-input').val() != '') {
    return false;
  }
  $('<div class="message loading new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure><span></span></div>').appendTo($('.mCSB_container'));
  updateScrollbar();

  setTimeout(function() {
    $('.message.loading').remove();
    $('<div class="message new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure>' + Fake[i] + '</div>').appendTo($('.mCSB_container')).addClass('new');
    setDate();
    updateScrollbar();
    i++;
  }, 1000 + (Math.random() * 20) * 100);

}