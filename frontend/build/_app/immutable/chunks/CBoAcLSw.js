import{s as o,g as c}from"./C4tR0XKQ.js";import{W as b,X as l,K as p,Q as a,m as d,t as _}from"./BxHNYr4i.js";let s=!1,i=Symbol();function m(e,n,r){const u=r[n]??(r[n]={store:null,source:p(void 0),unsubscribe:a});if(u.store!==e&&!(i in r))if(u.unsubscribe(),u.store=e,e==null)u.source.v=void 0,u.unsubscribe=a;else{var t=!0;u.unsubscribe=o(e,f=>{t?u.source.v=f:_(u.source,f)}),t=!1}return i in r?c(e):d(u.source)}function y(){const e={};function n(){b(()=>{for(var r in e)e[r].unsubscribe();l(e,i,{enumerable:!1,value:!0})})}return[e,n]}function N(e){var n=s;try{return s=!1,[e(),s]}finally{s=n}}export{m as a,N as c,y as s};
