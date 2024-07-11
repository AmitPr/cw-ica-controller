"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[6563],{1235:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>r,default:()=>h,frontMatter:()=>l,metadata:()=>i,toc:()=>a});var s=n(5893),o=n(1151);const l={title:"Go vs CosmWasm",sidebar_label:"Go vs CosmWasm",sidebar_position:2,slug:"/how-it-works/go-vs-cosmwasm"},r="Golang ICA Controller vs CosmWasm ICA Controller",i={id:"how-it-works/go-vs-cosmwasm",title:"Go vs CosmWasm",description:"The golang implementation of the ICS-27 controller standard is also widely deployed on IBC enabled chains, with some",source:"@site/docs/how-it-works/02-go-vs-cosmwasm.mdx",sourceDirName:"how-it-works",slug:"/how-it-works/go-vs-cosmwasm",permalink:"/cw-ica-controller/main/how-it-works/go-vs-cosmwasm",draft:!1,unlisted:!1,editUrl:"https://github.com/srdtrk/cw-ica-controller/tree/main/docs/docs/how-it-works/02-go-vs-cosmwasm.mdx",tags:[],version:"current",sidebarPosition:2,frontMatter:{title:"Go vs CosmWasm",sidebar_label:"Go vs CosmWasm",sidebar_position:2,slug:"/how-it-works/go-vs-cosmwasm"},sidebar:"docsSidebar",previous:{title:"Introduction",permalink:"/cw-ica-controller/main/how-it-works/introduction"},next:{title:"Channel Opening Handshake",permalink:"/cw-ica-controller/main/how-it-works/channel-handshake"}},c={},a=[];function d(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",p:"p",strong:"strong",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,o.a)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(t.h1,{id:"golang-ica-controller-vs-cosmwasm-ica-controller",children:"Golang ICA Controller vs CosmWasm ICA Controller"}),"\n",(0,s.jsx)(t.p,{children:"The golang implementation of the ICS-27 controller standard is also widely deployed on IBC enabled chains, with some\nnotable exceptions being the Cosmos Hub and Osmosis."}),"\n",(0,s.jsxs)(t.p,{children:["A contract developer could use the golang implementation of the ICA controller by submitting ",(0,s.jsx)(t.code,{children:"CosmosMsg::Stargate"}),"\nmessages which would be handled by the golang implementation of the ICA controller."]}),"\n",(0,s.jsx)(t.admonition,{type:"warning",children:(0,s.jsxs)(t.p,{children:["Using the golang implementation comes with a great disadvantage, which is that the golang implementation does not\nmake callbacks to the contract that submitted the ",(0,s.jsx)(t.code,{children:"CosmosMsg::Stargate"})," message! This means that the contract that\nsubmitted the ",(0,s.jsx)(t.code,{children:"CosmosMsg::Stargate"})," message will not be able to know the result of the ICA transaction, nor the\naddress of the newly created account on the other chain without intervention."]})}),"\n",(0,s.jsx)(t.admonition,{type:"tip",children:(0,s.jsxs)(t.p,{children:["Currently, Neutron and Nolus have custom bindings for the golang implementation of the ICA controller standard, which\nmake callbacks to the contract that submitted the ",(0,s.jsx)(t.code,{children:"CosmosMsg::Stargate"})," message. This means that you can use the\ngolang implementation of the ICA controller standard on Neutron and Nolus without any problems. However, any\napplications that are built on these chains will not be able to be ported to other chains."]})}),"\n",(0,s.jsxs)(t.p,{children:[(0,s.jsx)(t.code,{children:"cw-ica-controller"})," solves this problem by making callbacks to a contract that the developer specifies.\nThis means that the users of ",(0,s.jsx)(t.code,{children:"cw-ica-controller"})," can use the same contract on any chain that supports CosmWasm."]}),"\n",(0,s.jsxs)(t.p,{children:["But there is more! ",(0,s.jsx)(t.code,{children:"cw-ica-controller"})," is also able to do some things that the golang API limits you from doing."]}),"\n",(0,s.jsxs)(t.table,{children:[(0,s.jsx)(t.thead,{children:(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.th,{style:{textAlign:"center"},children:(0,s.jsx)(t.strong,{children:"Feature"})}),(0,s.jsx)(t.th,{style:{textAlign:"center"},children:(0,s.jsx)(t.code,{children:"cw-ica-controller"})}),(0,s.jsxs)(t.th,{style:{textAlign:"center"},children:["golang ",(0,s.jsx)(t.code,{children:"ica-controller"})]}),(0,s.jsx)(t.th,{style:{textAlign:"center"},children:"neutron bindings"}),(0,s.jsx)(t.th,{style:{textAlign:"center"},children:(0,s.jsx)(t.strong,{children:"Description"})})]})}),(0,s.jsxs)(t.tbody,{children:[(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Callbacks"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Golang implementation does not make callbacks"})]}),(0,s.jsxs)(t.tr,{children:[(0,s.jsxs)(t.td,{style:{textAlign:"center"},children:["Submit ",(0,s.jsx)(t.code,{children:"cosmwasm_std::CosmosMsg"}),"s"]}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsxs)(t.td,{style:{textAlign:"center"},children:["Golang implementation requires ICA transactions to be submitted in protobuf or ",(0,s.jsx)(t.code,{children:"proto3json"})," format."]})]}),(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Reopen an ICA channel with different params"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"In golang implementation, if the channel closes due to timeout the channel can only be reopened with the same parameters."})]}),(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Change owner"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsxs)(t.td,{style:{textAlign:"center"},children:[(0,s.jsx)(t.code,{children:"cw-ica-controller"})," uses ",(0,s.jsx)(t.a,{href:"https://github.com/larry0x/cw-plus-plus/tree/main/packages/ownable",children:(0,s.jsx)(t.code,{children:"cw-ownable"})})," for owner management"]})]}),(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Live upgrades"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"In golang implementation, new features require coordinated chain upgrades which could get blocked on upgrading CosmosSDK."})]}),(0,s.jsxs)(t.tr,{children:[(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Permanent channel closure"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u2705"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"\u274c"}),(0,s.jsx)(t.td,{style:{textAlign:"center"},children:"Golang implementation allows any relayer to always reopen any ICA channel (with the same parameters)."})]})]})]})]})}function h(e={}){const{wrapper:t}={...(0,o.a)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(d,{...e})}):d(e)}},1151:(e,t,n)=>{n.d(t,{Z:()=>i,a:()=>r});var s=n(7294);const o={},l=s.createContext(o);function r(e){const t=s.useContext(l);return s.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function i(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:r(e.components),s.createElement(l.Provider,{value:t},e.children)}}}]);