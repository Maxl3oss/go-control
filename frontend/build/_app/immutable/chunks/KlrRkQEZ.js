import{a as K,t as ne}from"./BIV2xHo9.js";import{aa as ve,M as he,aC as ie,U as y,a2 as R,a3 as pe,m as A,al as _e,am as W,a1 as M,V as I,aD as H,an as le,O as oe,ao as me,a8 as ce,q as ge,I as Y,aE as O,aF as J,aG as P,aH as we,a7 as be,aI as Ee,T as ye,Z as Te,H as xe,a4 as Se,ax as Ae,aJ as De,K as Ce,ad as X,ai as Ie,aK as Le,$ as He,p as ke,v as Ne,u as Z,x as z,a as Re,t as Me,w as Q,e as j,r as ee}from"./BxHNYr4i.js";import{s as Oe}from"./TTGbiMD1.js";import{a as Pe}from"./gUufbwpM.js";import{p as te}from"./CcGjO9XF.js";import{w as Ue}from"./C4tR0XKQ.js";function $e(t,e){return e}function Fe(t,e,a,i){for(var s=[],o=e.length,n=0;n<o;n++)we(e[n].e,s,!0);var c=o>0&&s.length===0&&a!==null;if(c){var u=a.parentNode;be(u),u.append(a),i.clear(),x(t,e[0].prev,e[o-1].next)}Ee(s,()=>{for(var p=0;p<o;p++){var v=e[p];c||(i.delete(v.k),x(t,v.prev,v.next)),ye(v.e,!c)}})}function qe(t,e,a,i,s,o=null){var n=t,c={flags:e,items:new Map,first:null},u=(e&ie)!==0;if(u){var p=t;n=y?R(Te(p)):p.appendChild(ve())}y&&pe();var v=null,_=!1,m=xe(()=>{var r=a();return Ie(r)?r:r==null?[]:ce(r)});he(()=>{var r=A(m),f=r.length;if(_&&f===0)return;_=f===0;let h=!1;if(y){var S=n.data===_e;S!==(f===0)&&(n=W(),R(n),M(!1),h=!0)}if(y){for(var g=null,w,b=0;b<f;b++){if(I.nodeType===8&&I.data===Se){n=I,h=!0,M(!1);break}var E=r[b],l=i(E,b);w=fe(I,c,g,null,E,l,b,s,e),c.items.set(l,w),g=w}f>0&&R(W())}if(!y){var d=Ae;Be(r,c,n,s,e,(d.f&H)!==0,i)}o!==null&&(f===0?v?le(v):v=oe(()=>o(n)):v!==null&&me(v,()=>{v=null})),h&&M(!0),A(m)}),y&&(n=I)}function Be(t,e,a,i,s,o,n,c){var F,q,B,V;var u=(s&De)!==0,p=(s&(O|P))!==0,v=t.length,_=e.items,m=e.first,r=m,f,h=null,S,g=[],w=[],b,E,l,d;if(u)for(d=0;d<v;d+=1)b=t[d],E=n(b,d),l=_.get(E),l!==void 0&&((F=l.a)==null||F.measure(),(S??(S=new Set)).add(l));for(d=0;d<v;d+=1){if(b=t[d],E=n(b,d),l=_.get(E),l===void 0){var ue=r?r.e.nodes_start:a;h=fe(ue,e,h,h===null?e.first:h.next,b,E,d,i,s),_.set(E,h),g=[],w=[],r=h.next;continue}if(p&&Ve(l,b,d,s),l.e.f&H&&(le(l.e),u&&((q=l.a)==null||q.unfix(),(S??(S=new Set)).delete(l))),l!==r){if(f!==void 0&&f.has(l)){if(g.length<w.length){var L=w[0],T;h=L.prev;var $=g[0],k=g[g.length-1];for(T=0;T<g.length;T+=1)ae(g[T],L,a);for(T=0;T<w.length;T+=1)f.delete(w[T]);x(e,$.prev,k.next),x(e,h,$),x(e,k,L),r=L,h=k,d-=1,g=[],w=[]}else f.delete(l),ae(l,r,a),x(e,l.prev,l.next),x(e,l,h===null?e.first:h.next),x(e,h,l),h=l;continue}for(g=[],w=[];r!==null&&r.k!==E;)(o||!(r.e.f&H))&&(f??(f=new Set)).add(r),w.push(r),r=r.next;if(r===null)continue;l=r}g.push(l),h=l,r=l.next}if(r!==null||f!==void 0){for(var C=f===void 0?[]:ce(f);r!==null;)(o||!(r.e.f&H))&&C.push(r),r=r.next;var N=C.length;if(N>0){var de=s&ie&&v===0?a:null;if(u){for(d=0;d<N;d+=1)(B=C[d].a)==null||B.measure();for(d=0;d<N;d+=1)(V=C[d].a)==null||V.fix()}Fe(e,C,de,_)}}u&&ge(()=>{var G;if(S!==void 0)for(l of S)(G=l.a)==null||G.apply()}),Y.first=e.first&&e.first.e,Y.last=h&&h.e}function Ve(t,e,a,i){i&O&&J(t.v,e),i&P?J(t.i,a):t.i=a}function fe(t,e,a,i,s,o,n,c,u,p){var v=(u&O)!==0,_=(u&Le)===0,m=v?_?Ce(s):X(s):s,r=u&P?X(n):n,f={i:r,v:m,k:o,a:null,e:null,prev:a,next:i};try{return f.e=oe(()=>c(t,m,r),y),f.e.prev=a&&a.e,f.e.next=i&&i.e,a===null?e.first=f:(a.next=f,a.e.next=f.e),i!==null&&(i.prev=f,i.e.prev=f.e),f}finally{}}function ae(t,e,a){for(var i=t.next?t.next.e.nodes_start:a,s=e?e.e.nodes_start:a,o=t.e.nodes_start;o!==i;){var n=He(o);s.before(o),o=n}}function x(t,e,a){e===null?t.first=a:(e.next=a,e.e.next=a&&a.e),a!==null&&(a.prev=e,a.e.prev=e&&e.e)}function at(t,e,a){var i=t.__className,s=Ge(e);y&&t.getAttribute("class")===s?t.__className=s:(i!==s||y&&t.getAttribute("class")!==s)&&(s===""?t.removeAttribute("class"):t.setAttribute("class",s),t.__className=s)}function Ge(t,e){return(t??"")+""}function re(t,e,a){if(a){if(t.classList.contains(e))return;t.classList.add(e)}else{if(!t.classList.contains(e))return;t.classList.remove(e)}}const U="http://43.229.149.77:8032/api";function Ke(t){return t.replace(/\u001b\[[0-9;]*m/g,"")}function We(t){return Ke(t.replace(/\\u([0-9a-fA-F]{4})/g,(e,a)=>String.fromCharCode(parseInt(a,16))))}function rt(t,e){const a=We(t),i=[];return a.split(`
`).forEach(o=>{const n=o.trim();if(!n)return;let c="info";n.includes("✓")&&(c="success"),n.includes("copied")&&(c="success"),n.includes("!")&&(c="warning"),n.includes("ignore")&&(c="warning"),n.includes("error")&&(c="error"),n.includes("$")&&(c="cmd"),i.push({id:e+1,message:n,type:c,timestamp:new Date})}),i}const Ye=async()=>{try{const t=await fetch(`${U}/get-site`);if(!t.ok)throw new Error(`HTTP error! Status: ${t.status}`);return(await t.json()).data}catch{return[]}},st=async t=>{try{const e=new FormData;e.append("file",t);const a=await fetch(`${U}/upload`,{method:"POST",body:e});if(!a.ok)throw new Error(`HTTP error! Status: ${a.status}`);return await a.json()}catch{return null}},D=async(t,e,a,i)=>{var s;try{const o=await fetch(`${U}/${e.toLowerCase()}/${t}`,{method:"POST",headers:{Accept:"text/event-stream"}});if(!o.ok)throw new Error(`HTTP error! Status: ${o.status}`);const n=(s=o.body)==null?void 0:s.getReader(),c=new TextDecoder;if(!n)throw new Error("ReadableStream not supported");let u="";for(;;){const{done:p,value:v}=await n.read();if(p)break;u+=c.decode(v,{stream:!0});const _=u.split(`
`);u=_.pop()||"";for(const m of _)if(m.startsWith("data:")){const r=m.replace("data:","").trim();a(r)}else!m.startsWith("event:")&&m.trim()!==""&&i(m.trim())}}catch(o){console.error("Error fetching data:",o)}},nt=async(t,e,a)=>D("npm-install",t,e,a),it=async(t,e,a)=>D("pull",t,e,a),lt=async(t,e,a)=>D("build",t,e,a),ot=async(t,e,a)=>D("stop-service",t,e,a),ct=async(t,e,a)=>D("deploy",t,e,a),ft=async(t,e,a)=>D("start-service",t,e,a);let se=Ue([]);var Je=ne('<a class="sites-item svelte-1ngiw0k"> </a>'),Xe=ne('<div class="h-screen fixed w-56 rounded-none bg-gray-800 border border-r-0 border-t-0 border-gray-700"><h3 class="p-1 pt-6 text-center text-lg font-medium text-gray-900 dark:text-white !rounded-b-none">Site</h3> <a href="/upload" class="sites-item svelte-1ngiw0k"><div class="w-full flex justify-center py-2 rounded-sm bg-orange-700">UpLoad</div></a> <!></div>');function ut(t,e){ke(e,!0);let a=Ne(te([]));const i=async()=>{const c=await Ye();se.set(c)};Z(()=>{i()}),Z(()=>se.subscribe(u=>{Me(a,te(u))}));var s=Xe(),o=Q(j(s),2),n=Q(o,2);qe(n,17,()=>A(a),$e,(c,u)=>{var p=Je(),v=j(p,!0);ee(p),z(_=>{Pe(p,"href",A(u).site_title),re(p,"site-active",e.siteTitle===A(u).site_title),Oe(v,_)},[()=>A(u).site_title.toLocaleUpperCase()]),K(c,p)}),ee(s),z(()=>re(o,"site-active",e.siteTitle==="form")),K(t,s),Re()}export{ut as H,lt as a,ct as b,ft as c,ot as d,qe as e,it as f,se as g,nt as h,$e as i,rt as p,at as s,re as t,st as u};
