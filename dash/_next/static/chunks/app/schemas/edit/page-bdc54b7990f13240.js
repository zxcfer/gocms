(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[598],{2071:function(e,l,s){Promise.resolve().then(s.bind(s,12428))},26723:function(e){e.exports=function(){var e=[],l=[],s={},n={},a={};function r(e){return"string"==typeof e?RegExp("^"+e+"$","i"):e}function i(e,l){return e===l?l:e===e.toLowerCase()?l.toLowerCase():e===e.toUpperCase()?l.toUpperCase():e[0]===e[0].toUpperCase()?l.charAt(0).toUpperCase()+l.substr(1).toLowerCase():l.toLowerCase()}function t(e,l,n){if(!e.length||s.hasOwnProperty(e))return l;for(var a=n.length;a--;){var r=n[a];if(r[0].test(l))return function(e,l){return e.replace(l[0],function(s,n){var a,r,t=(a=l[1],r=arguments,a.replace(/\$(\d{1,2})/g,function(e,l){return r[l]||""}));return""===s?i(e[n-1],t):i(s,t)})}(l,r)}return l}function o(e,l,s){return function(n){var a=n.toLowerCase();return l.hasOwnProperty(a)?i(n,a):e.hasOwnProperty(a)?i(n,e[a]):t(a,n,s)}}function d(e,l,s,n){return function(n){var a=n.toLowerCase();return!!l.hasOwnProperty(a)||!e.hasOwnProperty(a)&&t(a,a,s)===a}}function c(e,l,s){var n=1===l?c.singular(e):c.plural(e);return(s?l+" ":"")+n}return c.plural=o(a,n,e),c.isPlural=d(a,n,e),c.singular=o(n,a,l),c.isSingular=d(n,a,l),c.addPluralRule=function(l,s){e.push([r(l),s])},c.addSingularRule=function(e,s){l.push([r(e),s])},c.addUncountableRule=function(e){if("string"==typeof e){s[e.toLowerCase()]=!0;return}c.addPluralRule(e,"$0"),c.addSingularRule(e,"$0")},c.addIrregularRule=function(e,l){l=l.toLowerCase(),a[e=e.toLowerCase()]=l,n[l]=e},[["I","we"],["me","us"],["he","they"],["she","they"],["them","them"],["myself","ourselves"],["yourself","yourselves"],["itself","themselves"],["herself","themselves"],["himself","themselves"],["themself","themselves"],["is","are"],["was","were"],["has","have"],["this","these"],["that","those"],["echo","echoes"],["dingo","dingoes"],["volcano","volcanoes"],["tornado","tornadoes"],["torpedo","torpedoes"],["genus","genera"],["viscus","viscera"],["stigma","stigmata"],["stoma","stomata"],["dogma","dogmata"],["lemma","lemmata"],["schema","schemata"],["anathema","anathemata"],["ox","oxen"],["axe","axes"],["die","dice"],["yes","yeses"],["foot","feet"],["eave","eaves"],["goose","geese"],["tooth","teeth"],["quiz","quizzes"],["human","humans"],["proof","proofs"],["carve","carves"],["valve","valves"],["looey","looies"],["thief","thieves"],["groove","grooves"],["pickaxe","pickaxes"],["passerby","passersby"]].forEach(function(e){return c.addIrregularRule(e[0],e[1])}),[[/s?$/i,"s"],[/[^\u0000-\u007F]$/i,"$0"],[/([^aeiou]ese)$/i,"$1"],[/(ax|test)is$/i,"$1es"],[/(alias|[^aou]us|t[lm]as|gas|ris)$/i,"$1es"],[/(e[mn]u)s?$/i,"$1s"],[/([^l]ias|[aeiou]las|[ejzr]as|[iu]am)$/i,"$1"],[/(alumn|syllab|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$/i,"$1i"],[/(alumn|alg|vertebr)(?:a|ae)$/i,"$1ae"],[/(seraph|cherub)(?:im)?$/i,"$1im"],[/(her|at|gr)o$/i,"$1oes"],[/(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|automat|quor)(?:a|um)$/i,"$1a"],[/(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)(?:a|on)$/i,"$1a"],[/sis$/i,"ses"],[/(?:(kni|wi|li)fe|(ar|l|ea|eo|oa|hoo)f)$/i,"$1$2ves"],[/([^aeiouy]|qu)y$/i,"$1ies"],[/([^ch][ieo][ln])ey$/i,"$1ies"],[/(x|ch|ss|sh|zz)$/i,"$1es"],[/(matr|cod|mur|sil|vert|ind|append)(?:ix|ex)$/i,"$1ices"],[/\b((?:tit)?m|l)(?:ice|ouse)$/i,"$1ice"],[/(pe)(?:rson|ople)$/i,"$1ople"],[/(child)(?:ren)?$/i,"$1ren"],[/eaux$/i,"$0"],[/m[ae]n$/i,"men"],["thou","you"]].forEach(function(e){return c.addPluralRule(e[0],e[1])}),[[/s$/i,""],[/(ss)$/i,"$1"],[/(wi|kni|(?:after|half|high|low|mid|non|night|[^\w]|^)li)ves$/i,"$1fe"],[/(ar|(?:wo|[ae])l|[eo][ao])ves$/i,"$1f"],[/ies$/i,"y"],[/\b([pl]|zomb|(?:neck|cross)?t|coll|faer|food|gen|goon|group|lass|talk|goal|cut)ies$/i,"$1ie"],[/\b(mon|smil)ies$/i,"$1ey"],[/\b((?:tit)?m|l)ice$/i,"$1ouse"],[/(seraph|cherub)im$/i,"$1"],[/(x|ch|ss|sh|zz|tto|go|cho|alias|[^aou]us|t[lm]as|gas|(?:her|at|gr)o|[aeiou]ris)(?:es)?$/i,"$1"],[/(analy|diagno|parenthe|progno|synop|the|empha|cri|ne)(?:sis|ses)$/i,"$1sis"],[/(movie|twelve|abuse|e[mn]u)s$/i,"$1"],[/(test)(?:is|es)$/i,"$1is"],[/(alumn|syllab|vir|radi|nucle|fung|cact|stimul|termin|bacill|foc|uter|loc|strat)(?:us|i)$/i,"$1us"],[/(agend|addend|millenni|dat|extrem|bacteri|desiderat|strat|candelabr|errat|ov|symposi|curricul|quor)a$/i,"$1um"],[/(apheli|hyperbat|periheli|asyndet|noumen|phenomen|criteri|organ|prolegomen|hedr|automat)a$/i,"$1on"],[/(alumn|alg|vertebr)ae$/i,"$1a"],[/(cod|mur|sil|vert|ind)ices$/i,"$1ex"],[/(matr|append)ices$/i,"$1ix"],[/(pe)(rson|ople)$/i,"$1rson"],[/(child)ren$/i,"$1"],[/(eau)x?$/i,"$1"],[/men$/i,"man"]].forEach(function(e){return c.addSingularRule(e[0],e[1])}),["adulthood","advice","agenda","aid","aircraft","alcohol","ammo","analytics","anime","athletics","audio","bison","blood","bream","buffalo","butter","carp","cash","chassis","chess","clothing","cod","commerce","cooperation","corps","debris","diabetes","digestion","elk","energy","equipment","excretion","expertise","firmware","flounder","fun","gallows","garbage","graffiti","hardware","headquarters","health","herpes","highjinks","homework","housework","information","jeans","justice","kudos","labour","literature","machinery","mackerel","mail","media","mews","moose","music","mud","manga","news","only","personnel","pike","plankton","pliers","police","pollution","premises","rain","research","rice","salmon","scissors","series","sewage","shambles","shrimp","software","species","staff","swine","tennis","traffic","transportation","trout","tuna","wealth","welfare","whiting","wildebeest","wildlife","you",/pok[eé]mon$/i,/[^aeiou]ese$/i,/deer$/i,/fish$/i,/measles$/i,/o[iu]s$/i,/pox$/i,/sheep$/i].forEach(c.addUncountableRule),c}()},12428:function(e,l,s){"use strict";s.r(l),s.d(l,{default:function(){return D}});var n=s(57437),a=s(76540),r=s(34560),i=s(26723),t=s.n(i),o=s(21270),d=s(82670),c=s(57654),u=s(80244),m=s(50326),h=s(99497),x=s(2265),f=s(99773),j=s(32302),p=s(22782),v=s(18641),g=s(70094);let y=e=>{let{form:l}=e,{fields:s,append:a,remove:r}=(0,d.Dq)({name:"enums",control:l.control});return(0,n.jsxs)("div",{className:"space-y-4",children:[(0,n.jsx)("h3",{className:"text-md font-medium",children:"Enums"}),(0,n.jsx)(u.Wi,{control:l.control,name:"type",render:e=>{let{field:l}=e;return(0,n.jsx)(u.xJ,{className:"flex-shrink flex-grow flex-1 relative",children:(0,n.jsx)(u.zG,{})})}}),s.map((e,s)=>(0,n.jsxs)("div",{className:"flex flex-wrap items-stretch w-full relative gap-3",children:[(0,n.jsx)(u.Wi,{control:l.control,name:"enums.".concat(s,".value"),render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex-shrink flex-grow flex-1 relative",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{placeholder:"Enum value",...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"enums.".concat(s,".label"),render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex-shrink flex-grow flex-1 relative",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{placeholder:"Enum label",...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)("div",{className:"flex",children:(0,n.jsx)(c.z,{variant:"destructive",onClick:()=>{window.confirm("Are you sure you want to delete this enum?")&&r(s)},children:"Delete"})})]},e.value||s)),(0,n.jsxs)(c.z,{type:"button",variant:"outline",size:"sm",className:"mt-2",onClick:()=>a({value:"",label:""}),children:[(0,n.jsx)(g.Z,{size:16,className:"mr-2"}),"Add Enum value"]})]})};var b=s(96304),w=s(75006),N=s(86468),$=s(15e3),C=s(82628),I=s(95453);let _=e=>{let{form:l,editingField:s}=e,{data:r,isLoading:i,error:t}=(0,a.a)({queryKey:["schemas"],queryFn:b.cQ});return i?(0,n.jsx)(w.g,{}):t?(0,n.jsx)($.T,{error:t}):(0,n.jsxs)("div",{className:"space-y-4",children:[(0,n.jsx)(I.Z,{className:"my-4"}),(0,n.jsx)(u.Wi,{control:l.control,name:"type",render:e=>{let{field:l}=e;return(0,n.jsx)(u.xJ,{className:"flex-shrink flex-grow flex-1 relative",children:(0,n.jsx)(u.zG,{})})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"relation.type",render:e=>{let{field:l}=e;return(null==s?void 0:s.name)?(0,n.jsx)(p.I,{...l,readOnly:!0}):(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Relation type"}),(0,n.jsxs)(v.Ph,{onValueChange:l.onChange,defaultValue:l.value,name:"relation.type",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(v.i4,{children:(0,n.jsx)(v.ki,{placeholder:"Select a relation type"})})}),(0,n.jsxs)(v.Bw,{children:[(0,n.jsx)(v.Ql,{value:"o2o",children:"O2O"}),(0,n.jsx)(v.Ql,{value:"o2m",children:"O2M"}),(0,n.jsx)(v.Ql,{value:"m2m",children:"M2M"})]})]}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"relation.schema",render:e=>{let{field:l}=e;return(null==s?void 0:s.name)?(0,n.jsx)(p.I,{...l,readOnly:!0}):(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Relation schema"}),(0,n.jsxs)(v.Ph,{onValueChange:l.onChange,defaultValue:l.value,name:"relation.schema",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(v.i4,{children:(0,n.jsx)(v.ki,{placeholder:"Select a relation schema"})})}),(0,n.jsx)(v.Bw,{children:null==r?void 0:r.map(e=>(0,n.jsx)(v.Ql,{value:e.name,children:e.name},e.name))})]}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"relation.field",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:(0,n.jsx)(C.u,{icon:!0,tip:(0,n.jsxs)(n.Fragment,{children:[" ",(0,n.jsx)("p",{children:"The relation field is the field name of the relation schema that points back to the current editing field."}),(0,n.jsx)("p",{children:"For example:"}),(0,n.jsxs)("ul",{className:"list-decimal list-inside",children:[(0,n.jsxs)("li",{children:["The current editing field is ",(0,n.jsx)("code",{children:"post.post_meta"}),", and the relation schema is ",(0,n.jsx)("code",{children:"post_meta"})," (o2o), then the relation field is ",(0,n.jsx)("code",{children:"post_meta.post"})," (enter ",(0,n.jsx)("code",{children:"post"})," here)"]}),(0,n.jsxs)("li",{children:["The current editing field is ",(0,n.jsx)("code",{children:"post.comments"}),", and the relation schema is ",(0,n.jsx)("code",{children:"comment"})," (o2m), then the relation field is ",(0,n.jsx)("code",{children:"comment.post"})," (enter ",(0,n.jsx)("code",{children:"post"})," here)"]}),(0,n.jsxs)("li",{children:["The current editing field is ",(0,n.jsx)("code",{children:"post.categories"}),", and the relation schema is ",(0,n.jsx)("code",{children:"category"})," (m2m), then the relation field is ",(0,n.jsx)("code",{children:"category.posts"})," (enter ",(0,n.jsx)("code",{children:"posts"})," here)"]})]})]}),children:(0,n.jsx)("span",{className:"mr-1",children:"Relation Field"})})}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{autoComplete:"auto",...l,readOnly:!!(null==s?void 0:s.name)})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"relation.owner",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Owner"}),(0,n.jsx)(u.pf,{children:"Is the relation field the owner of the relation?"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"relation.owner",checked:l.value,onCheckedChange:l.onChange,disabled:!!(null==s?void 0:s.name),"aria-readonly":!!(null==s?void 0:s.name)})})]})}})]})};var k=s(53171);let S=e=>{var l;let{form:s,editingField:a}=e,r=s.watch("type"),i=s.watch("renderer.class"),t=(0,k.gc)(r,null!=a?a:f.bO),o=null!==(l=t.find(e=>e.class===i))&&void 0!==l?l:t[0];return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(u.Wi,{control:s.control,name:"server_name",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l,type:"hidden"})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:s.control,name:"type",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Type"}),(0,n.jsxs)(v.Ph,{onValueChange:l.onChange,defaultValue:l.value,name:"type",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(v.i4,{children:(0,n.jsx)(v.ki,{placeholder:"Select a type"})})}),(0,n.jsx)(v.Bw,{children:(0,n.jsxs)(j.x,{className:"h-72",children:[(0,n.jsxs)(v.DI,{children:[(0,n.jsx)(v.n5,{children:"Common"}),(0,n.jsx)(v.Ql,{value:"string",children:"Short text"}),(0,n.jsx)(v.Ql,{value:"text",children:"Long text"}),(0,n.jsx)(v.Ql,{value:"bool",children:"Boolean"}),(0,n.jsx)(v.Ql,{value:"int64",children:"Int"}),(0,n.jsx)(v.Ql,{value:"float64",children:"Float"}),(0,n.jsx)(v.Ql,{value:"media",children:"Media"}),(0,n.jsx)(v.Ql,{value:"relation",children:"Relation"})]}),(0,n.jsx)(v.U$,{}),(0,n.jsxs)(v.DI,{children:[(0,n.jsx)(v.n5,{children:"Complex"}),(0,n.jsx)(v.Ql,{value:"bytes",children:"Bytes"}),(0,n.jsx)(v.Ql,{value:"json",children:"Json"}),(0,n.jsx)(v.Ql,{value:"uuid",children:"UUID"})]}),(0,n.jsx)(v.U$,{}),(0,n.jsxs)(v.DI,{children:[(0,n.jsx)(v.n5,{children:"Advanced"}),(0,n.jsx)(v.Ql,{value:"enum",children:"Enum"}),(0,n.jsx)(v.Ql,{value:"time",children:"Time"})]})]})})]}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:s.control,name:"name",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Field name"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{autoComplete:"auto",placeholder:"age, address...",...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:s.control,name:"label",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Field label"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{placeholder:"Age, Address...",...l})}),(0,n.jsx)(u.zG,{})]})}}),"media"===r&&(0,n.jsx)(u.Wi,{control:s.control,name:"multiple",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Multiple"}),(0,n.jsx)(u.pf,{children:"Allow multiple values"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"unique",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}}),"enum"===r&&(0,n.jsx)(y,{form:s}),"relation"===r&&(0,n.jsx)(_,{form:s,editingField:a}),(0,n.jsx)(I.Z,{className:"my-4"}),(0,n.jsx)(u.Wi,{control:s.control,name:"renderer.class",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Renderer"}),(0,n.jsxs)(v.Ph,{onValueChange:l.onChange,defaultValue:l.value,name:"renderer.class",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(v.i4,{children:(0,n.jsx)(v.ki,{placeholder:"Select a renderer"})})}),(0,n.jsx)(v.Bw,{children:t.map(e=>(0,n.jsx)(v.Ql,{value:e.class,children:e.class},e.class))})]}),(0,n.jsx)(u.zG,{})]})}}),o?o.renderSettings(s):null]})},z=e=>{let{form:l}=e;return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(u.Wi,{control:l.control,name:"unique",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Unique"}),(0,n.jsx)(u.pf,{children:"Prevent duplicate values"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"unique",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"optional",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Optional"}),(0,n.jsx)(u.pf,{children:"Allow null values"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"optional",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"sortable",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Sortable"}),(0,n.jsx)(u.pf,{children:"Allow sorting"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"sortable",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"filterable",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Filterable"}),(0,n.jsx)(u.pf,{children:"Allow filtering"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"filterable",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"db.increment",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{className:"flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm",children:[(0,n.jsxs)("div",{className:"space-y-0.5",children:[(0,n.jsx)(u.lX,{children:"Increment"}),(0,n.jsx)(u.pf,{children:"Auto increment the value"})]}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(N.r,{name:"db.increment",checked:l.value,onCheckedChange:l.onChange,"aria-readonly":!0})})]})}})]})},F=e=>{let{form:l}=e,{register:s}=l;return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(u.Wi,{control:l.control,name:"default",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Default"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l,...s("default",{setValueAs:e=>""===e?null:e})})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"size",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Size"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{type:"number",...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"db.attr",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Attribute"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"db.collation",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Collation"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:l.control,name:"db.key",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:"Key"}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l})}),(0,n.jsx)(u.zG,{})]})}})]})};var q=s(97081),O=s(52235);function R(e){let{open:l,editingField:s,existingFields:a,onSave:r,onClose:i}=e,t=(0,d.cI)({resolver:(0,o.F)(f.gU),values:{...null!=s?s:f.bO},defaultValues:{...null!=s?s:f.bO}}),j=t.watch("name");(0,x.useEffect)(()=>{(null==s?void 0:s.name)||!j||t.setValue("label",(0,q.rV)(j))},[j]);let p=e=>{if(!(null==s?void 0:s.name)&&(null==a?void 0:a.find(l=>l.name===e.name))){t.setError("name",{type:"manual",message:"Field name already exists."});return}"relation"!==e.type&&delete e.relation,"enum"!==e.type&&delete e.enums,e.db&&0!==Object.keys(e.db).length||delete e.db,""===e.default&&delete e.default,e.optional||delete e.optional,null==r||r(e)};return(0,n.jsx)(m.yo,{open:l,onOpenChange:e=>{e||null==i||i()},children:(0,n.jsxs)(m.ue,{className:"lg:max-w-screen-lg overflow-y-auto max-h-screen py-0 w-full max-w-full field-edit-sheet sm:w-full sm:max-w-full md:w-3/4 md:max-w-3/4",onInteractOutside:e=>{e.preventDefault(),e.stopPropagation()},children:[(0,n.jsxs)(m.Tu,{className:"sticky top-0 z-10 bg-white py-5 text-left",children:[(0,n.jsx)(m.bC,{children:(null==s?void 0:s.name)?"Edit field: ".concat(s.name):"Create new field"}),(0,n.jsx)(m.Ei,{children:(null==s?void 0:s.name)?"Edit the field details below.":"Enter the field details below."}),(0,n.jsx)("button",{type:"button",onClick:i,className:"absolute top-2 right-2",children:(0,n.jsx)(O.Z,{size:20})})]}),(0,n.jsx)(u.l0,{...t,children:(0,n.jsx)("form",{onSubmit:t.handleSubmit(p),children:(0,n.jsx)("div",{className:"grid gap-4 py-4",children:(0,n.jsxs)(h.mQ,{defaultValue:"common",className:"w-full",children:[(0,n.jsxs)(h.dr,{className:"grid w-full grid-cols-3",children:[(0,n.jsx)(h.SP,{value:"common",children:"Common"}),(0,n.jsx)(h.SP,{value:"database",children:"Database"}),(0,n.jsx)(h.SP,{value:"advance",children:"Advance"})]}),(0,n.jsx)(h.nU,{value:"common",children:(0,n.jsx)("div",{className:"grid gap-4 py-4",children:(0,n.jsx)(S,{form:t,editingField:s})})}),(0,n.jsx)(h.nU,{value:"database",children:(0,n.jsx)("div",{className:"grid gap-4 py-4",children:(0,n.jsx)(F,{form:t})})}),(0,n.jsx)(h.nU,{value:"advance",children:(0,n.jsx)("div",{className:"grid gap-4 py-4",children:(0,n.jsx)(z,{form:t})})})]})})})}),(0,n.jsx)(m.FF,{className:"sticky bottom-0 z-10 bg-white py-5",children:(0,n.jsx)(c.z,{type:"submit",onClick:()=>{t.handleSubmit(p)()},children:"Save field"})})]})})}var E=s(27453),J=s(47907),W=s(12647);let X=e=>{var l,s,a;let{form:r,fields:i,editingSchema:t}=e;return(0,n.jsxs)(n.Fragment,{children:[r.formState.errors.fields&&(0,n.jsx)($.T,{description:null!==(a=r.formState.errors.fields.message)&&void 0!==a?a:null===(s=r.formState.errors.fields)||void 0===s?void 0:null===(l=s.root)||void 0===l?void 0:l.message}),(0,n.jsx)(u.Wi,{control:r.control,name:"name",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{className:"flex",children:(0,n.jsx)(C.u,{tip:"This is the name of your schema.",icon:!0,children:(0,n.jsx)("span",{className:"mr-1",children:"Name"})})}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l,autoComplete:"auto",placeholder:"Schema name"})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:r.control,name:"namespace",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{className:"flex",children:(0,n.jsx)(C.u,{icon:!0,tip:"This is the namespace of your schema. It will be used to generate the database table name and API endpoints.",children:(0,n.jsx)("span",{className:"mr-1",children:"Namespace"})})}),(0,n.jsx)(u.NI,{children:(0,n.jsx)(p.I,{...l,placeholder:"Schema namespace",className:(null==t?void 0:t.name)?"read-only:bg-gray-100":""})}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:r.control,name:"label_field",render:e=>{let{field:l}=e;return(0,n.jsxs)(u.xJ,{children:[(0,n.jsx)(u.lX,{children:(0,n.jsx)(C.u,{icon:!0,tip:"This is the namespace of your schema. It will be used to generate the database table name and API endpoints",children:(0,n.jsx)("span",{className:"mr-1",children:"Label Field"})})}),(0,n.jsxs)(v.Ph,{onValueChange:l.onChange,value:l.value,name:"label_field",children:[(0,n.jsx)(u.NI,{children:(0,n.jsx)(v.i4,{children:(0,n.jsx)(v.ki,{placeholder:"Select a label field"})})}),(0,n.jsx)(v.Bw,{children:(0,n.jsx)(j.x,{className:"h-72",children:i.map((e,l)=>(0,n.jsxs)(v.Ql,{value:e.name,children:[e.name," - ",e.label]},e.name))})})]}),(0,n.jsx)(u.zG,{})]})}}),(0,n.jsx)(u.Wi,{control:r.control,name:"disable_timestamp",render:e=>{let{field:l}=e;return(0,n.jsx)(u.xJ,{children:(0,n.jsx)(u.NI,{children:(0,n.jsxs)("div",{className:"flex items-center space-x-2",children:[(0,n.jsx)(N.r,{checked:l.value,onCheckedChange:l.onChange,name:"disable_timestamp",id:"disable_timestamp","aria-readonly":!0}),(0,n.jsx)(W._,{htmlFor:"disable_timestamp",children:"Disable timestamps"})]})})})}})]})};var P=s(5887),Z=s(77731),G=s(66806),V=s(33277);let Q=e=>{var l,s,a;let{form:r,fields:i,setEdittingField:t,removeField:o}=e,d=(0,n.jsx)(c.z,{type:"button",variant:"default",size:"sm",icon:(0,n.jsx)(g.Z,{size:16}),onClick:()=>t(f.bO),children:"New Field"});return(0,n.jsxs)("div",{className:"space-y-4",children:[r.formState.errors.fields&&(0,n.jsx)($.T,{description:null!==(a=r.formState.errors.fields.message)&&void 0!==a?a:null===(s=r.formState.errors.fields)||void 0===s?void 0:null===(l=s.root)||void 0===l?void 0:l.message}),d,(0,n.jsxs)(P.iA,{children:[(0,n.jsx)(P.xD,{children:(0,n.jsxs)(P.SC,{children:[(0,n.jsx)(P.ss,{children:"Label"}),(0,n.jsx)(P.ss,{children:"Name"}),(0,n.jsx)(P.ss,{children:"Type"}),(0,n.jsx)(P.ss,{children:"System"}),(0,n.jsx)(P.ss,{children:"Optional"}),(0,n.jsx)(P.ss,{className:"text-right",children:"Actions"})]})}),(0,n.jsx)(P.RM,{children:i.map((e,l)=>{var s;return(0,n.jsxs)(P.SC,{children:[(0,n.jsx)(P.pj,{className:"font-medium truncate",children:e.label}),(0,n.jsx)(P.pj,{className:"font-medium truncate",children:e.name}),(0,n.jsxs)(P.pj,{className:"truncate",children:[e.type,"relation"===e.type&&(0,n.jsx)(V.C,{variant:"secondary",className:"ml-2",children:null==e?void 0:null===(s=e.relation)||void 0===s?void 0:s.type})]}),(0,n.jsx)(P.pj,{children:e.is_system_field?(0,n.jsx)(V.C,{variant:"secondary",children:"System"}):null}),(0,n.jsx)(P.pj,{children:e.optional?null:(0,n.jsx)(V.C,{variant:"destructive",children:"Required"})}),(0,n.jsx)(P.pj,{className:"text-right",children:!e.is_system_field&&(0,n.jsxs)(n.Fragment,{children:[(0,n.jsxs)("button",{type:"button",className:"text-sm inline-flex flex-row items-center gap-1 hover:underline pr-3",onClick:()=>t(e),children:[(0,n.jsx)(Z.Z,{size:12})," Edit"]}),(0,n.jsxs)("button",{type:"button",className:"text-sm inline-flex flex-row items-center gap-1 text-red-800 hover:underline",onClick:()=>{window.confirm("Are you sure you want to delete this field?")&&o(l)},children:[(0,n.jsx)(G.Z,{size:12}),"Delete"]})]})})]},e.name)})})]}),i.length?d:null]})};var A=s(79984);let L=e=>{let l=(0,J.useRouter)(),{reloadAppConfig:s}=(0,x.useContext)(A.Il),{editingSchema:a}=e,[r,i]=(0,x.useState)(),m=(0,d.cI)({resolver:(0,o.F)(f.Pl),defaultValues:null!=a?a:f.IG,mode:"onChange"}),h=m.watch("name"),j=m.watch("label_field"),p=m.watch("fields"),{fields:v,append:g,update:y,remove:w}=(0,d.Dq)({name:"fields",control:m.control});(0,x.useEffect)(()=>{m.reset(null!=a?a:f.IG)},[null==a?void 0:a.name]),(0,x.useEffect)(()=>{(null==a?void 0:a.name)||!h||m.setValue("namespace",t()(h))},[h]),(0,x.useEffect)(()=>{if(p&&p.length>0&&!p.some(e=>e.name===j)){let e=p.find(e=>"string"===e.type);e&&m.setValue("label_field",e.name)}},[p,j]);let N=async e=>{let n=[],r=!1;for(let l of e.fields)l.server_name&&l.name!==l.server_name&&n.push({from:l.server_name,to:l.name}),l.name===e.label_field&&(r=!0);if(!r){E.h4.error("Label field is invalid."),m.setValue("label_field",""),m.setError("label_field",{type:"manual",message:"Label field is invalid."});return}try{var i;let r=await (0,b.uX)(e,null!==(i=null==a?void 0:a.name)&&void 0!==i?i:"",n);E.h4.success("Schema ".concat(r.name," saved successfully.")),l.push("/schemas/edit?schema=".concat(r.name)),m.reset(r),s()}catch(e){}};return(0,n.jsxs)("div",{className:"space-y-5",children:[(0,n.jsx)(u.l0,{...m,children:(0,n.jsxs)("form",{className:"grid gap-8 md:grid-cols-2 lg:grid-cols-3",onSubmit:m.handleSubmit(N),children:[(0,n.jsx)("div",{className:"relative flex-col items-start gap-8 md:flex",children:(0,n.jsxs)("fieldset",{className:"sticky top-5 w-full grid gap-5 rounded-lg border p-4",children:[(0,n.jsx)("legend",{className:"-ml-1 px-1 text-sm font-medium",children:"Schema"}),(0,n.jsx)(X,{form:m,fields:v,editingSchema:a}),(0,n.jsx)(c.z,{type:"submit",children:"Save"})]})}),(0,n.jsx)("div",{className:"relative flex lex-col rounded-xl lg:col-span-2",children:(0,n.jsxs)("fieldset",{className:"w-full grid gap-5 rounded-lg border p-4",children:[(0,n.jsx)("legend",{className:"-ml-1 px-1 text-sm font-medium",children:"Fields"}),(0,n.jsx)(Q,{form:m,fields:v,setEdittingField:i,removeField:w})]})})]})}),r?(0,n.jsx)(R,{open:!!r,editingField:r,existingFields:v,onClose:()=>i(void 0),onSave:e=>{if(null==r?void 0:r.name){let l=v.findIndex(e=>e.name===r.name);l>=0&&y(l,e)}else g(e);i(void 0)}}):null]})};function D(){let e=(0,J.useSearchParams)().get("schema"),{data:l,isLoading:s,error:i}=(0,a.a)({queryKey:["schema",e],queryFn:()=>(0,b.J1)(null!=e?e:""),retry:!1,refetchOnWindowFocus:!1});return((0,x.useEffect)(()=>{let s=(null==l?void 0:l.name)?"Edit schema: ".concat(l.name):"Create new schema",n=(null==l?void 0:l.name)?"Edit schema ".concat(l.name," to change the structure of your data."):"Create a new schema for your data.";return(0,r.D8)({title:s,description:n,breadcrumbs:[{name:"Schema",path:"/schemas"},{name:e?"Edit schema ".concat(e):"New schema",path:"/schemas/edit?schema="+e}]}),r.D8},[e,l]),s)?(0,n.jsx)(w.g,{}):i?(0,n.jsx)($.T,{error:i}):(0,n.jsx)(L,{editingSchema:l})}},99773:function(e,l,s){"use strict";s.d(l,{IG:function(){return t},Pl:function(){return r},bO:function(){return i},gU:function(){return a}});var n=s(30248);let a=n.Ry({name:n.Z_().trim().min(1,{message:"Field name is required"}),server_name:n.Z_().optional(),label:n.Z_().min(1,{message:"Field label is required"}),type:n.Km(["bool","time","json","uuid","bytes","enum","string","text","int","int8","int16","int32","int64","uint","uint8","uint16","uint32","uint64","float32","float64","relation","media"]),multiple:n.O7().optional(),size:n.oQ.number().optional(),unique:n.O7().optional(),optional:n.O7().optional(),default:n.Yj().nullable(),sortable:n.O7().optional(),filterable:n.O7().optional(),renderer:n.Ry({class:n.Z_().optional(),settings:n.IM(n.Z_(),n.Yj()).optional()}).optional(),enums:n.IX(n.Ry({value:n.Z_().min(1,{message:"Enum value is required"}),label:n.Z_().min(1,{message:"Enum label is required"})})).optional(),relation:n.Ry({schema:n.Z_(),field:n.Z_(),type:n.Km(["o2o","o2m","m2m"]),owner:n.O7().optional(),fk_columns:n.IM(n.Z_(),n.Z_()).optional().nullable(),junction_table:n.Z_().optional(),optional:n.O7().optional()}).optional(),db:n.Ry({attr:n.Z_().optional(),collation:n.Z_().optional(),increment:n.O7().optional(),key:n.Z_().optional()}).nullable().optional(),is_system_field:n.O7().optional()}).superRefine((e,l)=>{var s;return"enum"!==e.type||null!=e&&null!==(s=e.enums)&&void 0!==s&&!!s.length||l.addIssue({code:n.NL.custom,message:"Enums are required for enum type",path:["type"]})}).superRefine((e,l)=>{if("relation"===e.type){var s,a,r,i,t,o;(null==e?void 0:null===(s=e.relation)||void 0===s?void 0:s.type)&&(null==e?void 0:null===(a=e.relation)||void 0===a?void 0:a.schema)&&(null==e?void 0:null===(r=e.relation)||void 0===r?void 0:r.field)||l.addIssue({code:n.NL.custom,message:"Relation type, schema, field is required for relation type",path:["type"]}),(null==e?void 0:null===(i=e.relation)||void 0===i?void 0:i.type)||l.addIssue({code:n.NL.custom,message:"Relation type, schema, field is required for relation type",path:["relation.type"]}),(null==e?void 0:null===(t=e.relation)||void 0===t?void 0:t.schema)||l.addIssue({code:n.NL.custom,message:"Relation schema is required for relation type",path:["relation.schema"]}),(null==e?void 0:null===(o=e.relation)||void 0===o?void 0:o.field)||l.addIssue({code:n.NL.custom,message:"Relation field is required for relation type",path:["relation.field"]})}return!0}),r=n.Ry({name:n.Z_().trim().min(1,{message:"Schema name is required"}),namespace:n.Z_().trim().min(1,{message:"Schema namespace is required"}),label_field:n.Z_().trim().min(1,{message:"Schema label field is required"}),disable_timestamp:n.O7(),is_system_schema:n.O7().optional(),fields:n.IX(a).min(1,{message:"At least one field is required"})}).refine(e=>{var l;let s=null==e?void 0:null===(l=e.fields)||void 0===l?void 0:l.map(e=>e.name),n=[...new Set(s)];return s.length===n.length},{message:"Field names should be unique",path:["fields"]}).refine(e=>{var l;let s=null==e?void 0:null===(l=e.fields)||void 0===l?void 0:l.map(e=>e.label),n=[...new Set(s)];return s.length===n.length},{message:"Field labels should be unique",path:["fields"]}).refine(e=>{var l;return(null==e?void 0:null===(l=e.fields)||void 0===l?void 0:l.length)>0},{message:"At least one field is required",path:["fields"]}),i={name:"",server_name:"",label:"",type:"string",enums:[],default:"",size:0,multiple:!1,db:{attr:"",collation:"",increment:!1,key:""},renderer:{class:"",settings:{}},relation:{schema:"",field:"",type:"o2o",owner:!1,fk_columns:{},junction_table:"",optional:!1},unique:!1,optional:!0,sortable:!0,filterable:!1,is_system_field:!1},t={name:"",namespace:"",label_field:"",disable_timestamp:!1,fields:[]}},50326:function(e,l,s){"use strict";s.d(l,{Ei:function(){return v},FF:function(){return j},Tu:function(){return f},aM:function(){return c},bC:function(){return p},ue:function(){return x},yo:function(){return d}});var n=s(57437),a=s(2265),r=s(72936),i=s(62177),t=s(57742),o=s(22169);let d=r.fC,c=r.xz;r.x8;let u=r.h_,m=a.forwardRef((e,l)=>{let{className:s,...a}=e;return(0,n.jsx)(r.aV,{className:(0,o.cn)("fixed inset-0 z-50 bg-black/80  data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",s),...a,ref:l})});m.displayName=r.aV.displayName;let h=(0,t.j)("fixed z-50 gap-4 bg-background p-6 shadow-lg transition ease-in-out data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:duration-300 data-[state=open]:duration-500",{variants:{side:{top:"inset-x-0 top-0 border-b data-[state=closed]:slide-out-to-top data-[state=open]:slide-in-from-top",bottom:"inset-x-0 bottom-0 border-t data-[state=closed]:slide-out-to-bottom data-[state=open]:slide-in-from-bottom",left:"inset-y-0 left-0 h-full w-3/4 border-r data-[state=closed]:slide-out-to-left data-[state=open]:slide-in-from-left sm:max-w-sm",right:"inset-y-0 right-0 h-full w-3/4 border-l data-[state=closed]:slide-out-to-right data-[state=open]:slide-in-from-right sm:max-w-sm"}},defaultVariants:{side:"right"}}),x=a.forwardRef((e,l)=>{let{side:s="right",className:a,children:t,...d}=e;return(0,n.jsxs)(u,{children:[(0,n.jsx)(m,{}),(0,n.jsxs)(r.VY,{ref:l,className:(0,o.cn)(h({side:s}),a),...d,children:[t,(0,n.jsxs)(r.x8,{className:"absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-secondary",children:[(0,n.jsx)(i.Pxu,{className:"h-4 w-4"}),(0,n.jsx)("span",{className:"sr-only",children:"Close"})]})]})]})});x.displayName=r.VY.displayName;let f=e=>{let{className:l,...s}=e;return(0,n.jsx)("div",{className:(0,o.cn)("flex flex-col space-y-2 text-center sm:text-left",l),...s})};f.displayName="SheetHeader";let j=e=>{let{className:l,...s}=e;return(0,n.jsx)("div",{className:(0,o.cn)("flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",l),...s})};j.displayName="SheetFooter";let p=a.forwardRef((e,l)=>{let{className:s,...a}=e;return(0,n.jsx)(r.Dx,{ref:l,className:(0,o.cn)("text-lg font-semibold text-foreground",s),...a})});p.displayName=r.Dx.displayName;let v=a.forwardRef((e,l)=>{let{className:s,...a}=e;return(0,n.jsx)(r.dk,{ref:l,className:(0,o.cn)("text-sm text-muted-foreground",s),...a})});v.displayName=r.dk.displayName},96304:function(e,l,s){"use strict";s.d(l,{J1:function(){return i},cQ:function(){return a},uV:function(){return t},uX:function(){return r}});var n=s(84971);let a=async()=>(await (0,n.dX)("/schema")).filter(e=>!e.is_junction_schema),r=async(e,l,s)=>{var a,r;let i=null;return(null==(i=l?await (0,n.qb)("/schema/".concat(l),{schema:e,rename_fields:null!=s?s:[]}):await (0,n.SO)("/schema",e))?void 0:null===(a=i.fields)||void 0===a?void 0:a.length)&&(i.fields=(null!==(r=null==i?void 0:i.fields)&&void 0!==r?r:[]).map(e=>(e.server_name=e.name,e))),i},i=async e=>{var l,s;if(!e)return null;let a=await (0,n.dX)("/schema/".concat(e));return(null==a?void 0:null===(l=a.fields)||void 0===l?void 0:l.length)&&(a.fields=(null!==(s=null==a?void 0:a.fields)&&void 0!==s?s:[]).map(e=>(e.server_name=e.name,e))),a},t=e=>(0,n.HG)("/schema/".concat(e))}},function(e){e.O(0,[310,572,902,637,872,792,571,32,732,152,998,147,203,807,365,313,704,547,117,271,650,971,69,744],function(){return e(e.s=2071)}),_N_E=e.O()}]);