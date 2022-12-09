var gf = {
  setError: function (message) {
    $(".problem").show();
    $(".problem ul").empty();
    $(".problem ul").append("<li>" + message + "</li>");
  },
  setPrompt: function (message) {
    $(".prompt").show();
    $(".prompt ul").empty();
    $(".prompt ul").append("<li>" + message + "</li>");
  },
  setCookie: function (name, value, hours) {
    var d = new Date();
    d.setTime(d.getTime() + (hours * 60 * 60 * 1000));
    var expires = "expires=" + d.toGMTString();
    document.cookie = name + "=" + value + "; " + expires;
  },
  getCookie: function (cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
      var c = ca[i].trim();
      if (c.indexOf(name) === 0) return c.substring(name.length, c.length);
    }
    return "";
  },
  throttle: function (fn) {
    var timerId = null
    return function () {
      var arg = arguments[0]  //获取事件
      if (timerId) {
        return
      }
      timerId = setTimeout(function () {
        fn(arg)             //事件传入函数
        timerId = null
      }, 1000)
    }
  },

  getCaptcha: function () {

    successFun = function (data) {
      console.log(data);

      if (data.code === 0) {
        $('.captcha-img').html('<img src="' + data.data.captcha_base64 + '" />');
        $('input[name="captcha_id"]').val(data.data.captcha_id);
      }
    }

    $.ajax({
      url: '/captcha-get',
      type: 'get',
      dataType: 'json',
      success: successFun,
    });
  },
  // postByLink 通过a链接以form表单的提交的方式实现post请求,发送请求时，会携带一个参数，参数名为post_by_link，参数值为1,当后端返回响应时，发现此参数并且响应没有错误，则直接返回发起请求的页面

  // <a href="javascript:postByLink('/backend/1', {"name":"张三"})">提交</a>
  postByLink: function (action, data, confirm) {
    if (confirm !== "" && confirm !== undefined) {
      if (!window.confirm(confirm)) {
        return false;
      }
    }
    data['post_by_link'] = 1;
    var myForm = document.createElement("form");
    myForm.method = "post";
    myForm.action = action;
    for (var key in data) {
      var myInput = document.createElement("input");
      myInput.setAttribute("name", key);  // 为input对象设置name
      myInput.setAttribute("value", data[key]);  // 为input对象设置value
      myForm.appendChild(myInput);
    }


    document.body.appendChild(myForm);
    myForm.submit();
    document.body.removeChild(myForm);  // 提交后移除创建的form
  },
  initEditor: function (id) {
    var editor = new Quill('#' + id, {
      modules: {
        toolbar: [
          [{"font": []}, {"size": ["small", false, "large", "huge"]}], // custom dropdown

          ["bold", "italic", "underline", "strike"],

          [{"color": []}, {"background": []}],

          [{"script": "sub"}, {"script": "super"}],

          [{"header": 1}, {"header": 2}, "blockquote", "code-block"],

          [{"list": "ordered"}, {"list": "bullet"}, {"indent": "-1"}, {"indent": "+1"}],

          [{"direction": "rtl"}, {"align": []}],

          ["link", "image", "video", "formula"],

          ["clean"]
        ]
      },
      // placeholder: 'Free Write...',
      theme: 'snow',
    });

    let updateInProgress = false;

    editor.on('editor-change', () => {
      if (updateInProgress) return;
      var $images = $('#' + id + ' img')
      updateInProgress = true
      $images.each(function () {
        var imageSrc = $(this).attr('src')
        if (imageSrc && imageSrc[0] === 'd') {
          console.log('Starting image upload...')
          gf.uploadImageToImgurAndReplaceSrc($(this))
        }
      })
      updateInProgress = false

    })
    return editor;
  },
  uploadImageToImgurAndReplaceSrc: function ($image) {
    let base64 = $image.attr('src');

    let block = base64.split(";");
    // Get the content type of the image
    let contentType = block[0].split(":")[1];
    // get the real base64 content of the file
    let realData = block[1].split(",")[1];
    // Convert it to a blob to upload
    let blob = gf.b64toBlob(realData, contentType);
    // create form data
    const data = new FormData();
    // replace "file_upload" with whatever form field you expect the file to be uploaded to
    data.append('image', blob);
    $.ajax({
      url: '/posts-upload-image',
      type: 'post',
      data: data,
      processData: false,
      contentType: false,
      headers: {
        // Authorization: 'Client-ID ' + clientId
      },
      success: (response) => {
        $image.attr('src', response.data.url.replace(/^http(s?):/, ''));
      },
      error: (xhr, type, err) => {
        console.error(err)
        alert("Sorry we couldn't upload your image to Imgur.")
      }
    })
  },
  b64toBlob: function (b64Data, contentType, sliceSize) {
    contentType = contentType || '';
    sliceSize = sliceSize || 512;

    var byteCharacters = atob(b64Data);
    var byteArrays = [];

    for (var offset = 0; offset < byteCharacters.length; offset += sliceSize) {
      var slice = byteCharacters.slice(offset, offset + sliceSize);

      var byteNumbers = new Array(slice.length);
      for (var i = 0; i < slice.length; i++) {
        byteNumbers[i] = slice.charCodeAt(i);
      }

      var byteArray = new Uint8Array(byteNumbers);

      byteArrays.push(byteArray);
    }

    var blob = new Blob(byteArrays, {type: contentType});
    return blob;
  },

  wsInit: function (path) {
    let loc = window.location, url, protocol;
    if (loc.protocol === "https:") {
      protocol = "wss:";
    } else {
      protocol = "ws:";
    }

    url = `${protocol}//${loc.host}/${path}`;

    let wsApp = {}

    wsApp.init = function () {

      wsApp.maxReConnectAmount = 5;

      wsApp.currentReConnectAmount = 0;

      wsApp.lastHeartBeat = new Date().getTime();

      wsApp.overtime = 6000; //超时时间

      wsApp.send = function (d) { //发送消息

        d = JSON.stringify(d)

        if (!wsApp.ws || wsApp.ws.readyState !== 1) {
          console.log("WebSocket Server 初始化失败,无法发送消息");
          return
        }
        wsApp.ws.send(d)
      }

      wsApp.reConnect = function () { //重连
        clearInterval(wsApp.timer)
        console.log("socket 连接断开，第" + wsApp.currentReConnectAmount + "次重试");
        wsApp.init();
      }


      wsApp.checkConnect = function () { //检查是否需要重连

        if (wsApp.currentReConnectAmount > wsApp.maxReConnectAmount) {
          clearInterval(wsApp.timer)
          console.log("重试完成,无法连接,不再重试");
          return
        }

        wsApp.send(
          {
            type: "heart",
            data: {"ping": parseInt(new Date().getTime() / 1000).toString()},
          }
        ) //发送心跳
        if ((new Date().getTime() - wsApp.lastHeartBeat) > wsApp.overtime) {
          wsApp.currentReConnectAmount++;
          console.log("当前为第几次重试", wsApp.currentReConnectAmount);
          wsApp.reConnect();
        }
      }
      wsApp.ws = new WebSocket(url);
      try {
        // ws连接成功
        wsApp.ws.onopen = function () {
          console.log("WebSocket Server [" + url + "] 连接成功！");
        };
        // ws连接关闭
        wsApp.ws.onclose = function () {
          if (wsApp.ws) {
            wsApp.ws.close();
            wsApp.ws = null;
          }
          console.log("WebSocket Server [" + url + "] 连接关闭！");
        };
        // ws连接错误
        wsApp.ws.onerror = function () {
          if (wsApp.ws) {
            wsApp.ws.close();
            wsApp.ws = null;
          }
          console.log("WebSocket Server [" + url + "] 连接错误！");
        };
        // ws数据返回处理
        wsApp.ws.onmessage = function (result) {
          wsApp.lastHeartBeat = new Date().getTime();

        };

        if (wsApp.ws) {
          wsApp.timer = setInterval(wsApp.checkConnect, 5000);
        }
      } catch (e) {
        alert(e.message);
      }
    }

    wsApp.init()
    return wsApp
  }
}
