import{a as b,t as g}from"../chunks/BIV2xHo9.js";import"../chunks/B7fEFiRj.js";import{x as L,m as t,t as f,as as x,e as s,w as l,r,n as P}from"../chunks/BxHNYr4i.js";import{e as _}from"../chunks/Dn_gcW5b.js";import{i as F}from"../chunks/C36XFUe4.js";import{r as k,a as H}from"../chunks/gUufbwpM.js";import{a as S}from"../chunks/Ct7WMGyN.js";/* empty css                */function I(i){return function(...a){var o=a[0];return o.preventDefault(),i==null?void 0:i.apply(this,a)}}var R=g("<span>Hide</span>"),z=g("<span>Show</span>"),A=g(`<div class="flex items-center justify-center min-h-screen bg-gray-100 p-4"><div class="w-full max-w-md bg-white shadow-md rounded-lg p-6"><h2 class="text-2xl text-center mb-6">Login</h2> <form class="space-y-4"><div><label for="email" class="block text-sm font-medium mb-2">Email</label> <input type="email" id="email" placeholder="Enter your email" required class="w-full px-3 py-2 border rounded-md"></div> <div class="relative"><label for="password" class="block text-sm font-medium mb-2">Password</label> <input id="password" placeholder="Enter your password" required class="w-full px-3 py-2 border rounded-md pr-10"> <button type="button" class="absolute right-2 top-9 text-gray-500"><!></button></div> <div class="flex items-center justify-between"><div class="flex items-center"><input type="checkbox" id="remember-me" class="h-4 w-4 text-indigo-600 rounded"> <label for="remember-me" class="ml-2 text-sm">Remember me</label></div> <a href="" class="text-sm text-indigo-600 hover:text-indigo-500">Forgot password?</a></div> <button type="submit" class="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-indigo-700">Sign In</button> <div class="text-center mt-4"><p class="text-sm text-gray-600">Don't have an account? <a href="" class="text-indigo-600 hover:text-indigo-500">Sign up</a></p></div></form></div></div>`);function Q(i){let a=x(""),o=x(""),n=x(!1);function D(){console.log("Login attempt:",{email:t(a),password:t(o)})}var m=A(),h=s(m),p=l(s(h),2),u=s(p),w=l(s(u),2);k(w),r(u);var y=l(u,2),d=l(s(y),2);k(d);var c=l(d,2),E=s(c);{var j=e=>{var v=R();b(e,v)},q=e=>{var v=z();b(e,v)};F(E,e=>{t(n)?e(j):e(q,!1)})}r(c),r(y),P(6),r(p),r(h),r(m),L(()=>H(d,"type",t(n)?"text":"password")),S(w,()=>t(a),e=>f(a,e)),S(d,()=>t(o),e=>f(o,e)),_("click",c,()=>f(n,!t(n))),_("submit",p,I(D)),b(i,m)}export{Q as component};
