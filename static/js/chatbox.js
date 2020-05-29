$(document).ready(function () {
  $('#action_menu_btn').click(function () {
    $('.action_menu').toggle()
  })
})

const SenderType = {
  theirs: '',
  mine: '_send',
}

$(function () {
  const baseURI = '/chat/rooms/'
  const segments = window.location.pathname.split('/')
  if(segments.length < 1)
    return
  const roomId = segments[segments.length - 1]
  // $('#code').attr('src', `${baseURI}?string=${roomId}`)

  const content = (type, body, time) => {
    return `<div class="msg_cotainer${SenderType[type]}">
              ${body}
              <span class="msg_time${SenderType[type]}">${time}</span>
            </div>`;
  }

  const avatar = (imgUrl) => {
    return (
      `<div class="img_cont_msg">
        <img src="${imgUrl}" class="rounded-circle user_img_msg">
      </div>`
    );
  }

  const message = (type, body, time, imgUrl) => {
    var statEnd = 'start';
    type === SenderType.mine ? statEnd = 'end' : 'start'
    return `
          <div class="d-flex justify-content-${statEnd} mb-4">
           ${content(type, body, time)}
            ${avatar(imgUrl)}
          </div>
          `;
  }

  const uri = `${baseURI}/${roomId}/messages`
  $.get(uri, (data, status) => {
    console.log({ data, status });
  });
})
