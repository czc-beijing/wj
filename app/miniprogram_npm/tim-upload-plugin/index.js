module.exports = (function() {
var __MODS__ = {};
var __DEFINE__ = function(modId, func, req) { var m = { exports: {}, _tempexports: {} }; __MODS__[modId] = { status: 0, func: func, req: req, m: m }; };
var __REQUIRE__ = function(modId, source) { if(!__MODS__[modId]) return require(source); if(!__MODS__[modId].status) { var m = __MODS__[modId].m; m._exports = m._tempexports; var desp = Object.getOwnPropertyDescriptor(m, "exports"); if (desp && desp.configurable) Object.defineProperty(m, "exports", { set: function (val) { if(typeof val === "object" && val !== m._exports) { m._exports.__proto__ = val.__proto__; Object.keys(val).forEach(function (k) { m._exports[k] = val[k]; }); } m._tempexports = val }, get: function () { return m._tempexports; } }); __MODS__[modId].status = 1; __MODS__[modId].func(__MODS__[modId].req, m, m.exports); } return __MODS__[modId].m.exports; };
var __REQUIRE_WILDCARD__ = function(obj) { if(obj && obj.__esModule) { return obj; } else { var newObj = {}; if(obj != null) { for(var k in obj) { if (Object.prototype.hasOwnProperty.call(obj, k)) newObj[k] = obj[k]; } } newObj.default = obj; return newObj; } };
var __REQUIRE_DEFAULT__ = function(obj) { return obj && obj.__esModule ? obj.default : obj; };
__DEFINE__(1665834617297, function(require, module, exports) {
var global,factory;global=this,factory=function(){function e(t){return(e="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(t)}function t(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function r(e,t){for(var r=0;r<t.length;r++){var o=t[r];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}function o(e,t,o){return t&&r(e.prototype,t),o&&r(e,o),e}function n(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,o)}return r}function s(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){n(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,o,n=function(e,t){if(null==e)return{};var r,o,n={},a=Object.keys(e);for(o=0;o<a.length;o++)r=a[o],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)r=a[o],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var u="undefined"!=typeof wx&&"function"==typeof wx.getSystemInfoSync,f="undefined"!=typeof qq&&"function"==typeof qq.getSystemInfoSync,l="undefined"!=typeof tt&&"function"==typeof tt.getSystemInfoSync,c="undefined"!=typeof swan&&"function"==typeof swan.getSystemInfoSync,y="undefined"!=typeof my&&"function"==typeof my.getSystemInfoSync,d=u||f||l||c||y,p=f?qq:l?tt:c?swan:y?my:u?wx:{},h=function(t){if("object"!==e(t)||null===t)return!1;var r=Object.getPrototypeOf(t);if(null===r)return!0;for(var o=r;null!==Object.getPrototypeOf(o);)o=Object.getPrototypeOf(o);return r===o};function v(e){if(null==e)return!0;if("boolean"==typeof e)return!1;if("number"==typeof e)return 0===e;if("string"==typeof e)return 0===e.length;if("function"==typeof e)return 0===e.length;if(Array.isArray(e))return 0===e.length;if(e instanceof Error)return""===e.message;if(h(e)){for(var t in e)if(Object.prototype.hasOwnProperty.call(e,t))return!1;return!0}return!1}var m=function(){function e(){t(this,e),this.downloadUrl=""}return o(e,[{key:"request",value:function(e,t){var r=this;this.downloadUrl=e.downloadUrl||"";var o=(e.method||"PUT").toUpperCase(),n=e.url;if(e.qs){var a=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"&",r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"=";return v(e)?"":h(e)?Object.keys(e).map((function(o){var n=encodeURIComponent(o)+r;return Array.isArray(e[o])?e[o].map((function(e){return n+encodeURIComponent(e)})).join(t):n+encodeURIComponent(e[o])})).filter(Boolean).join(t):void 0}(e.qs);a&&(n+="".concat(-1===n.indexOf("?")?"?":"&").concat(a))}var s=new XMLHttpRequest;s.open(o,n,!0),s.responseType=e.dataType||"text";var i=e.headers||{};if(!v(i))for(var u in i)i.hasOwnProperty(u)&&"content-length"!==u.toLowerCase()&&"user-agent"!==u.toLowerCase()&&"origin"!==u.toLowerCase()&&"host"!==u.toLowerCase()&&s.setRequestHeader(u,i[u]);return s.onload=function(){t(null,r._xhrRes(s,r._xhrBody(s)))},s.onerror=function(e){var o=r._xhrBody(s);if(o)t(e,r._xhrRes(s,o));else{var n=s.statusText;n||0!==s.status||(n="CORS blocked or network error"),t(n,r._xhrRes(s,o))}},e.onProgress&&s.upload&&(s.upload.onprogress=function(t){var r=t.total,o=t.loaded,n=Math.floor(100*o/r);e.onProgress({total:r,loaded:o,percent:(n>=100?100:n)/100})}),s.send(e.resources),s}},{key:"_xhrRes",value:function(e,t){var r={};return e.getAllResponseHeaders().trim().split("\n").forEach((function(e){if(e){var t=e.indexOf(":"),o=e.substr(0,t).trim().toLowerCase(),n=e.substr(t+1).trim();r[o]=n}})),{statusCode:e.status,statusMessage:e.statusText,headers:r,data:t}}},{key:"_xhrBody",value:function(e){return 200===e.status&&e.responseURL&&this.downloadUrl?{location:this.downloadUrl}:{response:e.responseText}}}]),e}(),b=["unknown","image","video","audio","log"],g=function(){function e(){t(this,e),this.downloadUrl=""}return o(e,[{key:"request",value:function(e,t){var r=this,o=e.resources,n=void 0===o?"":o,a=e.headers,u=void 0===a?{}:a,f=e.url,l=e.downloadUrl,c=void 0===l?"":l;this.downloadUrl=c;var d=null,h="",v=c.match(/^(https?:\/\/[^/]+\/)([^/]*\/?)(.*)$/);h=(h=decodeURIComponent(v[3])).indexOf("?")>-1?h.split("?")[0]:h,u["Content-Type"]="multipart/form-data";var m={url:f,header:u,name:"file",filePath:n,formData:{key:h,success_action_status:200,"Content-Type":""},timeout:e.timeout||3e5};if(y){var g=m;g.name,m=s(s({},i(g,["name"])),{},{fileName:"file",fileType:b[e.fileType]})}return(d=p.uploadFile(s(s({},m),{},{success:function(e){r._handleResponse(e,t)},fail:function(e){r._handleResponse(e,t)}}))).onProgressUpdate((function(t){e.onProgress&&e.onProgress({total:t.totalBytesExpectedToSend,loaded:t.totalBytesSent,percent:Math.floor(t.progress)/100})})),d}},{key:"_handleResponse",value:function(e,t){var r=e.header,o={};if(r)for(var n in r)r.hasOwnProperty(n)&&(o[n.toLowerCase()]=r[n]);var a=+e.statusCode;200===a?t(null,{statusCode:a,headers:o,data:s(s({},e.data),{},{location:this.downloadUrl})}):t(e,{statusCode:a,headers:o,data:void 0})}}]),e}();return function(){function e(){t(this,e),this.retry=1,this.tryCount=0,this.systemClockOffset=0,this.httpRequest=d?new g:new m}return o(e,[{key:"uploadFile",value:function(e,t){var r=this;return this.httpRequest.request(e,(function(o,n){o&&r.tryCount<r.retry&&r.allowRetry(o)?(r.tryCount++,r.uploadFile(e,t)):(r.tryCount=0,t(o,n))}))}},{key:"allowRetry",value:function(e){var t=!1,r=!1;if(e){var o=e.headers&&(e.headers.date||e.headers.Date)||e.error&&e.error.ServerTime;try{var n=e.error&&e.error.Code,a=e.error&&e.error.Message;("RequestTimeTooSkewed"===n||"AccessDenied"===n&&"Request has expired"===a)&&(r=!0)}catch(u){}if(r&&o){var s=Date.now(),i=Date.parse(o);Math.abs(s+this.systemClockOffset-i)>=3e4&&(this.systemClockOffset=i-s,t=!0)}else 5===Math.floor(e.statusCode/100)&&(t=!0)}return t}}]),e}()},"object"==typeof exports&&"undefined"!=typeof module?module.exports=factory():"function"==typeof define&&define.amd?define(factory):(global=global||self).TIMUploadPlugin=factory();

}, function(modId) {var map = {}; return __REQUIRE__(map[modId], modId); })
return __REQUIRE__(1665834617297);
})()
//miniprogram-npm-outsideDeps=[]
//# sourceMappingURL=index.js.map