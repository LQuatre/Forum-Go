// const input = document.querySelector("#textarea");
// const messages = document.querySelector("#messages");
// const username = document.querySelector("#username");
// const send = document.querySelector("#send");

// const url = "ws://" + window.location.host + "/ws";
// const ws = new WebSocket(url);

// ws.onmessage = function (msg) {
//   insertMessage(JSON.parse(msg.data));
// };

// send.onclick = () => {
//   const message = {
//     username: username.value,
//     message: input.value,
//   };

//   console.log("message: " + JSON.stringify(message));
//   console.log("url: " + url);

//   ws.send(JSON.stringify(message));
//   input.value = "";
// };

// /**
//  * Insert a message into the UI
//  * @param {Message that will be displayed in the UI} messageObj
//  */
// function insertMessage(messageObj) {
//   // Create a div object which will hold the message
//   const message = document.createElement("div");

//   // Set the attribute of the message div
//   message.setAttribute("class", "chat-message");
//   console.log(
//     "name: " + messageObj.username + " content: " + messageObj.message
//   );
//   message.textContent = `${messageObj.username} : ${messageObj.message}`;

//   // Append the message to our chat div
//   messages.appendChild(message);

//   // Insert the message as the first message of our chat
//   messages.insertBefore(message, messages.firstChild);
// }

var $messages = $(".messages-content"),
  d,
  h,
  m,
  i = 0;

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

ws.onmessage = function (data) {
  wsMessage(JSON.parse(data));
};

$(window).on("load", function () {
  $messages.mCustomScrollbar();
  setTimeout(function () {
    // fakeMessage();
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
  setTimeout(function () {
    // fakeMessage();
  }, 1000 + Math.random() * 20 * 100);
}

$(".message-submit").click(function () {
  insertMessage();
  const message = {
    username: $("#username").val(),
    message: msg,
  };

  ws.send(JSON.stringify(message));
});

$(window).on("keydown", function (e) {
  if (e.which == 13) {
    insertMessage();
    return false;
  }
});

// var Fake = [
//   'Hi there, I\'m Fabio and you?',
//   'Nice to meet you',
//   'How are you?',
//   'Not too bad, thanks',
//   'What do you do?',
//   'That\'s awesome',
//   'Codepen is a nice place to stay',
//   'I think you\'re a nice person',
//   'Why do you think that?',
//   'Can you explain?',
//   'Anyway I\'ve gotta go now',
//   'It was a pleasure chat with you',
//   'Time to make a new codepen',
//   'Bye',
//   ':)'
// ]

function wsMessage(data) {
  if (data.username == $("#username").val()) {
    return;
  }
  $(
    '<div class="message loading new"><figure class="avatar"><img src="/static/img/jilt.png" /></figure><span></span></div>'
  ).appendTo($(".mCSB_container"));
  updateScrollbar();

  setTimeout(function () {
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
    i++;
  }, 1000 + Math.random() * 20 * 100);
}
