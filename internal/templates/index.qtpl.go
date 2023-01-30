// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamIndex(qw422016 *qt422016.Writer, profiles []FullProfileView) {
	qw422016.N().S(`
<!DOCTYPE html>
<html class="page-root" lang="uk">

<head>
    <title>GitHub profile views statistic badge</title>
    <meta name="description" content="GitHub profile views statistic badge">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="icon" sizes="192x192" href="/assets/files/icon192.png">
    <link rel="shortcut icon" sizes="16x16" href="/assets/files/favicon.ico">
    <link rel="shortcut icon" type="image/png" sizes="32x32" href="/assets/files/icon192.png">

    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&display=swap" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/apexcharts"></script>

    <style>html{box-sizing:border-box}*,:before,:after{box-sizing:inherit}html{-webkit-text-size-adjust:100%;-moz-text-size-adjust:100%;text-size-adjust:100%}body{-webkit-overflow-scrolling:touch}html,body,div,span,applet,object,iframe,h1,h2,h3,h4,h5,h6,p,blockquote,pre,a,abbr,acronym,address,big,cite,code,del,dfn,em,img,ins,kbd,q,s,samp,small,strike,strong,sub,sup,tt,var,b,u,i,center,dl,dt,dd,ol,ul,li,fieldset,form,label,legend,table,caption,tbody,tfoot,thead,tr,th,td,article,aside,canvas,details,embed,figure,figcaption,footer,header,menu,nav,output,ruby,section,summary,time,mark,audio,video{margin:0;padding:0;border:0}article,aside,details,figcaption,figure,footer,header,main,menu,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block}audio:not([controls]){display:none;height:0}a{background-color:transparent}abbr[title]{border-bottom:none;text-decoration:underline;-webkit-text-decoration:underline dotted;text-decoration:underline dotted}b,strong{font-weight:bolder}dfn{font-style:italic}mark{background-color:#ff0;color:#000}svg:not(:root){overflow:hidden}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}hr{box-sizing:content-box;height:0;overflow:visible}button,input,select,textarea{font:inherit;margin:0}button,input{overflow:visible}button,select{text-transform:none}button,[type=button],[type=reset],[type=submit]{-webkit-appearance:button;-moz-appearance:button;appearance:button}input,textarea,button,select,a{-webkit-tap-highlight-color:transparent}address{font-style:normal}a:focus:not(:focus-visible),select:focus:not(:focus-visible),input:focus:not(:focus-visible),textarea:focus:not(:focus-visible){outline:0}button::-moz-focus-inner,[type=button]::-moz-focus-inner,[type=reset]::-moz-focus-inner,[type=submit]::-moz-focus-inner{border-style:none;padding:0}button,input[type=reset],input[type=button],input[type=submit]{cursor:pointer}button[disabled],input[disabled]{cursor:default}button{-webkit-appearance:none;-moz-appearance:none;appearance:none;background:0 0;padding:0;border:0;border-radius:0;line-height:1}button:focus:not(:focus-visible){outline:0}a,a:hover{text-decoration:none}[href="javascript:void();"],[href="javascript:"]{cursor:default}ul,ol{list-style:none}blockquote,q{quotes:none}blockquote:before,blockquote:after,q:before,q:after{content:none}table{border-collapse:collapse;border-spacing:0}input[type=text],input[type=password],input[type=date],input[type=datetime],input[type=datetime-local],input[type=email],input[type=month],input[type=number],input[type=search],input[type=tel],input[type=time],input[type=url],input[type=week],textarea{box-sizing:border-box}[type=checkbox],[type=radio]{box-sizing:border-box;margin:0;padding:0}input[type=search]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield;outline-offset:-2px}input[type=search]::-webkit-search-decoration,input[type=search]::-webkit-search-cancel-button{-webkit-appearance:none;appearance:none}input[type=number]::-webkit-inner-spin-button,input[type=number]::-webkit-outer-spin-button{height:auto;-webkit-appearance:none;appearance:none}::-webkit-file-upload-button{-webkit-appearance:button;appearance:button;font:inherit}input[type=number]{-webkit-appearance:textfield;-moz-appearance:textfield;appearance:textfield}select{width:100%;height:20px;border:0;background:0 0}textarea{resize:none;border:0;overflow:auto}::-webkit-input-placeholder{color:#777;line-height:normal}::-moz-placeholder{color:#777;line-height:normal}::placeholder{color:#777;line-height:normal}[hidden]{display:none}.headline{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-weight:500;color:#111028;line-height:120%}.headline--lvl1{font-size:44px}.headline--lvl2{font-size:38px}.headline--lvl3{font-size:28px}.headline--lvl4{font-size:18px}.button{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:18px;font-weight:500;display:-webkit-inline-flex;display:inline-flex;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center;height:48px;border-radius:4px;padding-left:32px;padding-right:32px;border-width:2px;border-style:solid;border-color:transparent;transition:background-color .225s ease,box-shadow .225s ease,border-color .225s ease,color .225s ease}.button--black{color:#fff;background-color:#111028;border-color:#111028}.button--black:hover{background-color:#3f3a92;border-color:#3f3a92}.button--black:active{background-color:#2f2c6d;border-color:#2f2c6d}.button--black:focus-visible{background-color:#3f3a92;border-color:#dcdbf0;box-shadow:none}.button--black:disabled{color:#79769c;background-color:#dcdbf0;border-color:#dcdbf0}.input{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column}.input.is-active .input__clear{display:-webkit-flex;display:flex}.input.is-error .input__error{display:-webkit-flex;display:flex}.input.is-error .input__element{border-color:#e9331a}.input.is-error .input__error-text{display:block}.input__label{font-size:16px;line-height:19px;color:#8c8ba3}.input__wrapper{position:relative;margin-top:8px}.input__element{font-size:18px;line-height:22px;color:#111028;height:48px;width:100%;padding-left:20px;padding-right:48px;border-radius:4px;border-width:1px;border-style:solid;border-color:transparent;background-color:#f3f3f6;transition:border-color .225s ease,box-shadow .225s ease}.input__element:hover{border-color:#918fc7;box-shadow:0 0 6px rgba(67,66,88,.16)}.input__element:focus{box-shadow:none;outline:0}.input__element::-webkit-input-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::-moz-placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element::placeholder{font-size:18px;line-height:22px;color:#adacc2}.input__element:disabled{pointer-events:none}.input__element:disabled::-webkit-input-placeholder{color:#d4d4d4}.input__element:disabled::-moz-placeholder{color:#d4d4d4}.input__element:disabled::placeholder{color:#d4d4d4}.input__clear{display:none;position:absolute;right:20px;top:12px;width:24px;height:24px;-webkit-align-items:center;align-items:center;-webkit-justify-content:center;justify-content:center}.input__error{display:none;position:absolute;right:20px;top:12px}.input__error-text{display:none;font-size:14px;line-height:17px;color:#e9331a;position:absolute;top:calc(100% + 8px);left:0}.select__label{font-size:16px;line-height:120%;color:#8c8ba3}.select__wrapper{height:40px;padding-top:8px;padding-bottom:8px;padding-left:15px;padding-right:15px;border-radius:4px;border:1px solid #dcdbf0;margin-top:12px;transition:border-color .225s ease}.select__wrapper:focus-within{border-color:#b6b5d9}.select__element{font-size:18px;line-height:25px;color:#201f3a;cursor:pointer}.select__element option{font-size:16px;color:#7c7b88}.select__element option:disabled{color:#cbcadb}.select__element:focus-visible{outline:0}.checkbox{font-size:18px;color:#3c3a59;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;height:24px;cursor:pointer}.checkbox__input{position:absolute;-webkit-appearance:none;-moz-appearance:none;appearance:none}.checkbox__input:focus-visible{width:18px;height:18px;outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.checkbox__input:checked+.checkbox__element{background-color:#3c3a59;background-size:17px 14px;background-repeat:no-repeat;background-position:3px center;background-image:url("data:image/svg+xml,%3Csvg width='17' height='14' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath fill-rule='evenodd' clip-rule='evenodd' d='M16.173.26a1 1 0 0 1 .067 1.413L5.534 13.449.293 8.207a1 1 0 1 1 1.414-1.414l3.759 3.758L14.76.327A1 1 0 0 1 16.173.26Z' fill='%23fff'/%3E%3C/svg%3E")}.checkbox__element{width:24px;height:24px;border:1px solid #3c3a59;border-radius:4px;background-color:#fff;margin-right:12px;transition:background-color .225s ease}.filters__group{padding-top:32px}.filters__sub-headline{font-size:18px;font-weight:600;color:#3c3a59;height:24px;line-height:24px}.filters__visibility{width:28px;height:28px;position:relative;cursor:pointer}.filters__visibility:focus-visible{outline:0;box-shadow:0 0 0 2px #000,0 0 0 3px #fff;border-radius:4px}.filters__visibility:before{content:"";position:absolute;top:50%;left:50%;width:14px;height:2px;margin-left:-7px;margin-top:-1px;background-color:#828282}.filters__elements{margin-top:16px;max-height:160px;overflow:auto}.filters__elements--no-scroll{overflow:initial}.filters__elements::-webkit-scrollbar{width:4px}.filters__elements::-webkit-scrollbar-track{border-radius:8px;background-color:#dddce4}.filters__elements::-webkit-scrollbar-thumb{border-radius:8px;background-color:#b8b6e2}.filters__filled-elements{display:-webkit-flex;display:flex;-webkit-flex-wrap:wrap;flex-wrap:wrap;gap:8px;margin-top:16px}.filters__element:not(:first-child){margin-top:16px}html{height:100%}body{font-family:"Inter",arial,helvetica neue,helvetica,sans-serif;font-size:16px;color:#21212b;display:grid;grid-template-rows:auto 1fr auto;min-height:100%;background-color:#f2f2f3}.wrapper{max-width:1248px;width:100%;padding-left:16px;padding-right:16px;margin-left:auto;margin-right:auto}.disabled-scroll{overflow-y:hidden}.visually-hidden:not(:focus):not(:active){position:absolute;width:1px;height:1px;margin:-1px;border:0;padding:0;white-space:nowrap;-webkit-clip-path:inset(100%);clip-path:inset(100%);clip:rect(0 0 0 0);overflow:hidden}.header__wrapper{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.header__search{width:100%;padding:40px;border-radius:24px;background-color:#fff;margin-top:80px}.header__action-group{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;-webkit-justify-content:space-between;justify-content:space-between;margin-top:32px}.header__input{margin-top:32px}.footer__group{color:rgba(0,0,0,.8);display:grid;grid-template-columns:1fr auto auto;-webkit-align-items:start;align-items:start;padding-top:40px;padding-bottom:60px;padding-left:0;padding-right:0;border-top:1px solid #d0cfdf}.footer__copyrights{font-size:14px;line-height:20px;color:#585764;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;-webkit-column-gap:14px;-moz-column-gap:14px;column-gap:14px}.footer__dev{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;margin-right:95px}.footer__dev:last-child{margin-right:45px}.footer__dev-label{font-size:12px;line-height:17px;color:#9f9ead}.footer__dev-name{font-size:14px;line-height:20px;color:#585764;margin-top:8px}.footer__dev-link{font-size:14px;line-height:20px;color:#3c2df9;-webkit-align-self:flex-start;align-self:flex-start;margin-top:4px}</style>
    <style>.footer{background:#020111}.footer__wrapper{padding:80px 0 32px;display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between}.footer__info{color:#fff}.footer__title{display:block;font-weight:600;font-size:24px;line-height:120%;margin-bottom:16px;color:#fff}.footer__subtitle{line-height:140%;margin-bottom:48px}.footer__map{margin-bottom:56px}.footer__copyrights{display:-webkit-flex;display:flex;gap:14px;color:#cfd2d7;font-size:14px}.footer__support{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:16px;-webkit-align-items:flex-start;align-items:flex-start}.footer__link{padding:16.5px 32px 16.5px 40px;border:1px solid #636077;border-radius:4px;color:#fff;transition:border .3s}.footer__link:hover{border:1px solid #fff}.footer__link:active{border:1px solid #fff;background:#060423}.footer__figure{display:-webkit-flex;display:flex;gap:12px}.header{box-shadow:0 4px 12px rgba(134,132,177,.1);background:#fff;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;min-height:80px}.header__wrapper{display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;-webkit-align-items:center;align-items:center;width:1216px;height:100%}.header__logo{line-height:0;display:block;margin-right:auto}.header__profile{position:relative;margin-left:40px}.modal.active{opacity:1;pointer-events:all}.modal{position:absolute;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;width:400px;right:-57%;top:166%;border-radius:10px;background:#fff;opacity:0;pointer-events:none}.modal::after{content:"";position:absolute;right:40px;top:-12px;width:0;height:0;border-left:12px solid transparent;border-right:12px solid transparent;border-bottom:12px solid #fff;border-radius:1px}.modal__profile{display:-webkit-flex;display:flex;gap:16px;padding:24px;border-bottom:1px solid #dddce4}.modal__profile-name{font-weight:600;font-size:18px;line-height:140%;color:#020111}.modal__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:4px;font-size:14px;line-height:140%;color:#0d1dab}.modal__log-out{padding:24px}.modal__button{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:11px}.hero{background-image:url(/assets/files/bg.jpg);background-repeat:no-repeat;background-size:cover;background-position:center;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;color:#fff;margin:64px 0 104px;height:600px;border-radius:16px;padding:0 300px;text-align:center}.hero__title{font-size:64px;line-height:120%;font-weight:600;margin-bottom:24px}.hero__subtitle{font-size:21px;line-height:140%;margin-bottom:48px}.hero__button{display:-webkit-flex;display:flex;padding:16px 56px;background:#fff;border-radius:4px;transition:.9s}.hero__figure{display:-webkit-flex;display:flex;gap:10px;-webkit-align-items:center;align-items:center;color:#24292f}.hero__button:hover figcaption{color:#3f3a92}.hero__button:hover path{fill:#3f3a92}.hero__button:active figcaption{color:#2f2c6d}.hero__button:active path{fill:#2f2c6d}.history{margin-bottom:104px}.history__title{font-size:48px;line-height:120%;margin-bottom:16px}.history__subtitle{line-height:140%;color:#7c7b88;width:590px;margin-bottom:40px}.history__main{background:#fff;border-radius:16px;padding:40px}.history__main-values{width:100%;display:-webkit-flex;display:flex;-webkit-justify-content:space-between;justify-content:space-between;margin-bottom:12px;color:#7c7b88}.history__list{height:717px;overflow:auto;width:100%}.history__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;padding:24px 0;border-top:1px solid #dddce4;margin-right:24px}.history__user{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:16px;width:480px;margin-right:16px}.history__user-name{font-weight:600;font-size:21px;line-height:140%;color:#24292f}.history__link{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:4px;color:#0d1dab}.history__user-github{width:400px;text-overflow:ellipsis;white-space:nowrap;overflow:hidden}.history__badges{display:-webkit-flex;display:flex;gap:2px;font-size:14px;line-height:140%}.history__badge{display:-webkit-flex;display:flex;color:#fff}.history__badge-title{display:block;padding:2px 6px 3px 8px;background:#000;border:1px solid #e4eaf1;border-radius:4px 0 0 4px}.history__badge-count{display:block;padding:2px 6px 3px 8px;background:#6d96ff;border-width:1px 1px 1px 0;border-style:solid;border-color:#e4eaf1;border-radius:0 4px 4px 0}.history__user-time{margin-left:auto;color:#7c7b88}.history ::-webkit-scrollbar{width:10px}.history ::-webkit-scrollbar-bottom{width:4px}.history ::-webkit-scrollbar-track{background-color:#f2f2f8}.history ::-webkit-scrollbar-thumb{background-color:#b8b6e2;border-radius:10px}.stat-reg{margin-top:64px}.stat-reg__title{font-weight:600;font-size:48px;line-height:120%;margin-bottom:16px}.stat-reg__subtitle{line-height:140%;color:#7c7b88;margin-bottom:40px;max-width:524px}.stat-reg .sorting__list{text-align:center;top:116%;right:3%;bottom:unset!important;left:unset!important}.stat-reg .sorting__current{margin-left:auto}.profile{margin-top:64px;color:#111028}.profile__card{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center;gap:32px;margin-bottom:40px}.profile__info{display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;gap:8px}.profile__name{font-weight:600;font-size:48px;line-height:120%;color:#020111}.profile__link{display:block;font-size:18px;line-height:140%;color:#0d1dab;margin-bottom:12px}.profile__statistics{background:#fff;border-radius:16px;width:100%;height:488px;padding:40px;margin-bottom:104px;color:#111028}.profile__header{display:-webkit-flex;display:flex;-webkit-align-items:flex-end;align-items:flex-end;margin-bottom:40px}.profile__title{font-weight:600;font-size:28px;line-height:120%}.profile__view-count{font-weight:400;font-size:18px;line-height:140%;margin-left:auto;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.profile__circle{display:inline-block;width:16px;height:16px;border-radius:50%;background:#24292f;margin-right:12px}.profile__select{margin-left:64px;width:90px}.sorting__list{text-align:center;top:116%;right:3%;bottom:unset!important;left:unset!important}.sorting__current{margin-left:auto}.badge{margin-bottom:131px}.badge__title{font-weight:600;font-size:48px;line-height:120%;margin-bottom:16px}.badge__subtitle{line-height:140%;color:#7c7b88;margin-bottom:40px}.badge__example{background:#fff;border-radius:16px;width:100%;word-wrap:break-word;padding:32px}.badge__code{margin-bottom:16px;line-height:140%}.badge__copy{position:relative;margin-left:auto;overflow:hidden}.badge__copy-button{display:-webkit-flex;display:flex;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;gap:16px;width:200px;height:45px;margin-left:auto;background:#111028;color:#fff;font-weight:600;font-size:18px;line-height:140%;border-radius:4px;transition:background .3s}.badge__copy-button:hover{background:#232149}.badge__copy-button:active{background:#060423}.badge__copy-info{display:-webkit-flex;display:flex;-webkit-justify-content:center;justify-content:center;-webkit-align-items:center;align-items:center;gap:16px;position:absolute;right:-275px;top:0;width:275px;pointer-events:none;background:#3f3a92;color:#fff;height:45px;font-weight:600;font-size:18px;line-height:140%;border-radius:4px}@-webkit-keyframes anim{0%{right:-275px}10%{right:0}90%{right:0}}@keyframes anim{0%{right:-275px}10%{right:0}90%{right:0}}.wrapper-u8{max-width:1216px;margin:0 auto;color:#020111}body{font-size:18px}main{background:#f2f2f8}.sorting{margin:40px 0 24px;padding-bottom:16px;border-bottom:1px solid #dddce4;font-size:18px;display:-webkit-flex;display:flex}.sorting__item{display:-webkit-flex;display:flex;margin-right:32px;position:relative}.sorting__item h4{display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__item:last-child{margin-right:0}.sorting__title{padding:8px;padding-right:12px}.sorting__select{position:relative;display:-webkit-flex;display:flex;-webkit-align-items:center;align-items:center}.sorting__current{width:-webkit-max-content;width:-moz-max-content;width:max-content;border:0;font-weight:600;font-size:18px;line-height:140%;color:#111028;padding:0;padding-right:8px;text-align:left;background:0 0;position:relative}.sorting__icon{padding:7px 4px}.sorting__list{opacity:0;pointer-events:none;display:-webkit-flex;display:flex;-webkit-flex-direction:column;flex-direction:column;padding:8px 0;border:1px solid #ededf8;box-shadow:0 6px 16px rgba(17,16,40,.15);border-radius:4px;background:#fff;color:#3c3a59;font-size:18px;line-height:140%;position:absolute;bottom:-156px;left:80px;z-index:1;width:-webkit-max-content;width:-moz-max-content;width:max-content}.sorting__option{padding:8px 16px;text-align:left;background:inherit;transition:background .3s;color:#7c7b88}.sorting__option:hover{background:#f2f2f8}.sorting__list.is-visible{transition:opacity .4s;opacity:1;pointer-events:visible}</style>
</head>

<body>
    <header class="header">
	<div class="header__wrapper wrapper-u8">
		<a href="/" class="header__logo">
			<img class="header__logo-img" src="/assets/files/logo.svg" alt=""></a>
		<div class="header__stars">
			<iframe src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
				frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
		</div>
	</div>
</header>
    <main class="main">

<div class="wrapper-u8">
    <section class="hero">
        <h1 class="hero__title">GitHub profile view statistics</h1>
        <p class="hero__subtitle">Get a badge for your GitHub profile and detailed view statistics for a selected time
            period</p>
        <a href="/login/github" class="hero__button">
            <figure class="hero__figure">
                <svg width="26" height="26" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path
                        d="M13 1C6.37346 1 1 6.5085 1 13.3038C1 18.74 4.43837 23.3519 9.20639 24.9789C9.80609 25.0927 10.0263 24.712 10.0263 24.387C10.0263 24.0936 10.0151 23.1243 10.01 22.0962C6.67152 22.8405 5.96707 20.6445 5.96707 20.6445C5.42121 19.2224 4.6347 18.8443 4.6347 18.8443C3.54598 18.0806 4.71676 18.0963 4.71676 18.0963C5.9218 18.1831 6.55632 19.3642 6.55632 19.3642C7.62659 21.2452 9.36356 20.7014 10.0483 20.3871C10.156 19.5918 10.4671 19.0491 10.8102 18.7418C8.14488 18.4306 5.34291 17.3756 5.34291 12.6612C5.34291 11.318 5.81169 10.2203 6.57938 9.35871C6.45477 9.04876 6.04406 7.7974 6.69561 6.10263C6.69561 6.10263 7.70329 5.77194 9.99648 7.36384C10.9536 7.09114 11.9802 6.9545 13 6.94987C14.0199 6.9545 15.0472 7.09114 16.0062 7.36384C18.2967 5.77194 19.303 6.10263 19.303 6.10263C19.9561 7.7974 19.5452 9.04876 19.4206 9.35871C20.19 10.2203 20.6556 11.3179 20.6556 12.6612C20.6556 17.3868 17.8483 18.4274 15.1763 18.732C15.6066 19.1138 15.9902 19.8626 15.9902 21.0105C15.9902 22.6567 15.9762 23.9817 15.9762 24.387C15.9762 24.7144 16.1922 25.098 16.8006 24.9772C21.566 23.3485 25 18.7382 25 13.3038C25 6.5085 19.6273 1 13 1Z"
                        fill="#24292F" />
                    <path
                        d="M5.49426 18.527C5.46791 18.5881 5.37398 18.6064 5.28862 18.5645C5.20156 18.5244 5.15262 18.441 5.18086 18.3796C5.20674 18.3167 5.30066 18.2992 5.38753 18.3414C5.47478 18.3815 5.52447 18.4656 5.49426 18.527ZM6.08454 19.067C6.02732 19.1214 5.91542 19.0961 5.83947 19.0102C5.76098 18.9244 5.7463 18.8097 5.80436 18.7544C5.86337 18.7001 5.97189 18.7255 6.05057 18.8114C6.12906 18.8981 6.1443 19.0121 6.08445 19.0671M6.48952 19.7579C6.41592 19.8103 6.29564 19.7612 6.22138 19.6518C6.14788 19.5424 6.14788 19.4112 6.22298 19.3586C6.29752 19.306 6.41592 19.3533 6.49121 19.4619C6.56462 19.5731 6.56462 19.7044 6.48942 19.758M7.17429 20.5582C7.10851 20.6325 6.96846 20.6126 6.86588 20.5111C6.76104 20.4119 6.73177 20.2711 6.79774 20.1967C6.86428 20.1222 7.00517 20.1431 7.10851 20.2438C7.21269 20.3428 7.2445 20.4846 7.17439 20.5582M8.05934 20.8283C8.03045 20.9246 7.89549 20.9685 7.75959 20.9275C7.62388 20.8854 7.53503 20.7725 7.56242 20.6751C7.59065 20.5781 7.72618 20.5325 7.86311 20.5763C7.99864 20.6183 8.08767 20.7303 8.05944 20.8283M9.06674 20.9429C9.07013 21.0444 8.95484 21.1285 8.81216 21.1304C8.66864 21.1336 8.5526 21.0514 8.55109 20.9517C8.55109 20.8492 8.66375 20.7658 8.80718 20.7634C8.94985 20.7605 9.06674 20.842 9.06674 20.9429ZM10.0563 20.904C10.0734 21.003 9.97419 21.1047 9.83255 21.1317C9.69326 21.1578 9.56432 21.0967 9.54654 20.9986C9.52922 20.897 9.6303 20.7954 9.76931 20.7691C9.91123 20.7438 10.0382 20.8033 10.0563 20.904Z"
                        fill="#24292F" />
                </svg>

                <figcaption class="hero__figcaption">Sign up with GitHub</figcaption>
            </figure>
        </a>
    </section>

    <section class="history">
        <h2 class="history__title">Registration history</h2>
        <p class="history__subtitle">Recent registered users who posted a badge on their GitHub profile</p>
        <div class="history__main">
            <div class="history__main-values">
                <span>User:</span>
                <span>Time, data:</span>
            </div>
            <ul class="history__list">
                `)
	for _, profile := range profiles {
		qw422016.N().S(`
                <li class="history__card">
                    <div class="history__user">
                        <img src="https://avatars.githubusercontent.com/u/`)
		qw422016.E().S(profile.SocialProviderUserID)
		qw422016.N().S(`?v=4&s=56" alt="`)
		qw422016.E().S(profile.Username)
		qw422016.N().S(`" class="history__user-picture">
                        <div class="history__user-info">
                            <div class="history__user-name">`)
		qw422016.E().S(profile.GetName())
		qw422016.N().S(`</div>
                            <a href="https://github.com/`)
		qw422016.E().S(profile.Username)
		qw422016.N().S(`" class="history__link">
                                <img src="/assets/files/link.svg" width="16" height="16" alt="link">
                                <span class="history__user-github">github.com/`)
		qw422016.E().S(profile.Username)
		qw422016.N().S(`</span>
                            </a>
                        </div>
                    </div>
                    <div class="history__badges">
                        <div class="history__badge">
                            <span class="history__badge-title">Views per day</span>
                            <span class="history__badge-count">`)
		qw422016.N().DL(profile.DayCount)
		qw422016.N().S(`</span>
                        </div>
                        <div class="history__badge">
                            <span class="history__badge-title">per week</span>
                            <span class="history__badge-count">`)
		qw422016.N().DL(profile.WeekCount)
		qw422016.N().S(`</span>
                        </div>
                        <div class="history__badge">
                            <span class="history__badge-title">per month</span>
                            <span class="history__badge-count">`)
		qw422016.N().DL(profile.MonthCount)
		qw422016.N().S(`</span>
                        </div>
                    </div>
                    <div class="history__user-time">
                        `)
		qw422016.E().S(profile.CreatedAt.Format("15:04, Jan _2 2006"))
		qw422016.N().S(`
                    </div>
                </li>
                `)
	}
	qw422016.N().S(`
            </ul>
        </div>
    </section>
</div>

</main>
`)
	streamfooter(qw422016)
	qw422016.N().S(`
</body>
</html>
`)
}

func WriteIndex(qq422016 qtio422016.Writer, profiles []FullProfileView) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamIndex(qw422016, profiles)
	qt422016.ReleaseWriter(qw422016)
}

func Index(profiles []FullProfileView) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteIndex(qb422016, profiles)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
