(this["webpackJsonpportal-ui"]=this["webpackJsonpportal-ui"]||[]).push([[162],{955:function(e,t,c){"use strict";c.r(t);var a=c(15),s=c(3),l=c(0),n=c(38),i=c(417),r=c(418),o=c(419),d=c(316),j=c(474),b=c(463),u=c(330),O=c(50),h=c(30),p=c(250),x=c(116),f=c(328),y=c(122),v=c.n(y),k=c(260),m=c(2),g={setErrorSnackMessage:h.e},C=Object(n.b)(null,g);t.default=Object(k.a)((function(e){return Object(p.a)(Object(s.a)({},x.h))}))(C((function(e){var t=e.classes,c=e.closeDeleteModalAndRefresh,s=e.deleteOpen,n=e.selectedUser,h=e.setErrorSnackMessage,p=Object(l.useState)(!1),x=Object(a.a)(p,2),y=x[0],k=x[1];return null===n?Object(m.jsx)("div",{}):Object(m.jsxs)(i.a,{open:s,onClose:function(){c(!1)},classes:t,className:t.root,"aria-labelledby":"alert-dialog-title","aria-describedby":"alert-dialog-description",children:[Object(m.jsxs)(r.a,{id:"alert-dialog-title",className:t.title,children:[Object(m.jsx)("div",{className:t.titleText,children:"Delete User"}),Object(m.jsx)("div",{className:t.closeContainer,children:Object(m.jsx)(f.a,{"aria-label":"close",className:t.closeButton,onClick:function(){c(!0)},disableRipple:!0,size:"small",children:Object(m.jsx)(v.a,{})})})]}),Object(m.jsxs)(o.a,{children:[y&&Object(m.jsx)(d.a,{}),Object(m.jsxs)(j.a,{id:"alert-dialog-description",children:["Are you sure you want to delete user ",Object(m.jsx)("br",{}),Object(m.jsx)("b",{children:n.accessKey}),"?"]})]}),Object(m.jsxs)(b.a,{children:[Object(m.jsx)(u.a,{type:"button",variant:"outlined",onClick:function(){c(!1)},color:"primary",disabled:y,children:"Cancel"}),Object(m.jsx)(u.a,{onClick:function(){y||null!==n&&(k(!0),O.a.invoke("DELETE","/api/v1/user?name=".concat(encodeURI(n.accessKey)),{id:n.id}).then((function(e){k(!1),c(!0)})).catch((function(e){k(!1),h(e)})))},type:"button",variant:"outlined",color:"secondary",autoFocus:!0,children:"Delete"})]})]})})))}}]);
//# sourceMappingURL=162.69652a57.chunk.js.map