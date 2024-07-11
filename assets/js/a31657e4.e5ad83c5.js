"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[4274],{9052:(e,o,t)=>{t.r(o),t.d(o,{assets:()=>l,contentTitle:()=>r,default:()=>p,frontMatter:()=>c,metadata:()=>a,toc:()=>d});var n=t(5893),i=t(1151),s=t(7326);const c={title:"Introduction",sidebar_label:"Introduction",sidebar_position:0,slug:"/"},r="CosmWasm ICA Controller",a={id:"intro",title:"Introduction",description:"Welcome to the documentation for CosmWasm Interchain Accounts Controller. This document will guide you through",source:"@site/versioned_docs/version-v0.5.x/00-intro.mdx",sourceDirName:".",slug:"/",permalink:"/cw-ica-controller/v0.5/",draft:!1,unlisted:!1,editUrl:"https://github.com/srdtrk/cw-ica-controller/tree/main/docs/versioned_docs/version-v0.5.x/00-intro.mdx",tags:[],version:"v0.5.x",sidebarPosition:0,frontMatter:{title:"Introduction",sidebar_label:"Introduction",sidebar_position:0,slug:"/"},sidebar:"docsSidebar",next:{title:"Overview",permalink:"/cw-ica-controller/v0.5/contract-api/intro"}},l={},d=[{value:"High Level Overview",id:"high-level-overview",level:2}];function h(e){const o={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",img:"img",p:"p",...(0,i.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(o.h1,{id:"cosmwasm-ica-controller",children:"CosmWasm ICA Controller"}),"\n",(0,n.jsx)(s.Z,{type:"concepts"}),"\n",(0,n.jsx)(s.Z,{type:"basics"}),"\n",(0,n.jsxs)(o.p,{children:["Welcome to the documentation for CosmWasm Interchain Accounts Controller. This document will guide you through\nunderstanding the ",(0,n.jsx)(o.a,{href:"https://github.com/cosmos/ibc/tree/main/spec/app/ics-027-interchain-accounts",children:"ICS-27"}),"\nInterchain Accounts protocol and how to use ",(0,n.jsx)(o.code,{children:"cw-ica-controller"})," to create and manage interchain accounts on\nany IBC enabled CosmWasm chain."]}),"\n",(0,n.jsx)(o.p,{children:"The CosmWasm ICA Controller is a CosmWasm contract that implements the ICS-27 interchain accounts controller in\npure Rust. It is designed to be used by other CosmWasm contracts to create and manage interchain accounts on\nthe chain where the contract is deployed."}),"\n",(0,n.jsx)(o.h2,{id:"high-level-overview",children:"High Level Overview"}),"\n",(0,n.jsxs)(o.p,{children:["The following diagram shows how ",(0,n.jsx)(o.code,{children:"cw-ica-controller"})," works at a high level."]}),"\n",(0,n.jsx)(o.p,{children:(0,n.jsx)(o.img,{alt:"High Level Overview",src:t(4358).Z+"",width:"1127",height:"441"})}),"\n",(0,n.jsxs)(o.p,{children:["The ",(0,n.jsx)(o.code,{children:"cw-ica-controller"})," contract code is deployed on a chain that supports IBC CosmWasm. This chain does not need\nto support ICS-27 interchain accounts nor does it need to support any custom IBC bindings. Then when an external\naccount or a contract instantiates a ",(0,n.jsx)(o.code,{children:"cw-ica-controller"})," contract, the contract will initiate the ICS-27 handshake\nwith a chain that supports ICS-27 interchain accounts based on the options provided by the caller."]}),"\n",(0,n.jsx)(o.admonition,{type:"note",children:(0,n.jsxs)(o.p,{children:["The counterparty chain need not be a CosmWasm chain. It can be any chain that uses ibc-go and supports ",(0,n.jsx)(o.code,{children:"ICS-27"}),".\nSuch as CosmosHub, Osmosis, etc."]})}),"\n",(0,n.jsxs)(o.p,{children:["Then the rest of the ICS-27 handshake is completed by the relayers automatically. Both the hermes relayer and the\ngo relayer support ",(0,n.jsx)(o.code,{children:"ICS-27"})," interchain accounts. Once the handshake is complete, the ",(0,n.jsx)(o.code,{children:"cw-ica-controller"})," contract\nmakes a callback to the callback contract if one was provided during instantiation."]})]})}function p(e={}){const{wrapper:o}={...(0,i.a)(),...e.components};return o?(0,n.jsx)(o,{...e,children:(0,n.jsx)(h,{...e})}):h(e)}},7326:(e,o,t)=>{t.d(o,{Z:()=>m});var n=t(7294),i=t(394),s=t(512),c=t(8388);function r(){for(var e=arguments.length,o=new Array(e),t=0;t<e;t++)o[t]=arguments[t];return(0,c.m6)((0,s.W)(o))}var a=t(5893);const l=i.fC,d=i.xz,h=n.forwardRef(((e,o)=>{let{className:t,align:n="center",sideOffset:s=4,...c}=e;return(0,a.jsx)(i.VY,{ref:o,align:n,sideOffset:s,className:r("z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",t),...c})}));h.displayName=i.VY.displayName;const p={concepts:{color:"#54ffe0",label:"Concepts",isBright:!0,description:"Learn about the concepts behind 'cw-ica-controller'"},basics:{color:"#F69900",label:"Basics",isBright:!0,description:"Learn the basics of 'cw-ica-controller'"},"ibc-go":{color:"#ff1717",label:"IBC-Go",description:"This section includes IBC-Go specific content"},cosmjs:{color:"#6836D0",label:"CosmJS",description:"This section includes CosmJS specific content"},cosmwasm:{color:"#05BDFC",label:"CosmWasm",description:"This section includes CosmWasm specific content"},protocol:{color:"#00B067",label:"Protocol",description:"This section includes content about protocol specifcations"},advanced:{color:"#f7f199",label:"Advanced",isBright:!0,description:"The content in this section is for advanced users researching"},developer:{color:"#AABAFF",label:"Developer",isBright:!0,description:"This section includes content for external developers using the 'cw-ica-controller'"},tutorial:{color:"#F46800",label:"Tutorial",description:"This section includes a tutorial"},"guided-coding":{color:"#F24CF4",label:"Guided Coding",description:"This section includes guided coding"}},m=e=>{let{type:o,version:t}=e;const n=p[o]||p["ibc-go"],i=n.description||"";return(0,a.jsxs)(l,{children:[(0,a.jsx)(d,{children:(0,a.jsxs)("span",{style:{backgroundColor:n.color,borderRadius:"2px",color:n.isBright?"black":"white",padding:"0.3rem",marginBottom:"1rem",marginRight:"0.25rem",display:"inline-block"},children:[n.label,t?` ${t}`:""]})}),(0,a.jsx)(h,{children:i})]})}},4358:(e,o,t)=>{t.d(o,{Z:()=>n});const n=t.p+"assets/images/cw-ica-controller-02461d2650f1678266617056d26d355d.svg"}}]);