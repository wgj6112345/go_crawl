package baidu

import (
	"fmt"
	"regexp"
	"testing"
)

var body = `<div class="result-op c-container xpath-log"  srcid="1577"  id="12" tpl="xueshu_detail" mu="http://xueshu.baidu.com/s?wd=paperuri:(1v2r00k0pt260tj0ss2w0pa0wy042721)&filter=sc_long_sign&sc_ks_para=q%3D%E5%B8%82%E5%9F%9F%E7%A4%BE%E4%BC%9A%E6%B2%BB%E7%90%86%E7%8E%B0%E4%BB%A3%E5%8C%96%E8%B7%AF%E5%BE%84%E7%A0%94%E7%A9%B6%E2%80%94%E2%80%94%E4%BB%A5%E7%BB%8D%E5%85%B4%E5%B8%82%E4%B8%BA%E4%BE%8B&tn=SE_baiduxueshu_c1gjeupa&ie=utf-8&sc_us=12010585807122850015" data-op="{'y':'FCFE577F'}" data-click="{'p1':'12','rsv_bdr':'0','fm':'alop',rsv_stl:''}">

            <h3 class="t">
                                                                            <a href="http://www.baidu.com/link?url=D6ukhURUB3dVbDtt8u6f_8Z9NCIuZskWX8eLruTJcjbTfFJJH0L-LiZSfLkwStBdAv_CWEUt9BsSKlpc_jI9W5PoLTMceKhrZBLcbS6rPga5s_xGKyy1PNte9dUyBVW_42HLaEBuH-yNC_A5G4HAjtMoMQaAr5v3XEUqeiZ6dWrNMQCCuwYkyhbX1w2Hfg83EK9a9vgt0DjbtYiTpNWcuP1PalgVRGHstKFw4J2FxG_6W5Cebg06LAZRB6eFAgV8heQmndVIhrqBr8CsWRYMScZQXXhx6yq46qDh2DWjHajZExtzEo3l67VCAugrsPQlixcH5QfcCjpEyQ1eb7bBZjF5mTf8nkOZbdziLWBZhBvtVoCFOlDk2GbOZPNYm7WVcZOLqTQl_Sh3RI6BMdy_po__0s89EI0uMjZJt81Ri37SE8pZ_t0ryeuKsLoBT5R-uwuCzbYEAfSngynpRs7ctA5UD_KOlcbp4cnAo_nbjZe3WM6SwL0_nZMNDrMW9TMm" target="_blank">   
                                        <em>市域社会治理现代化</em>路径研究——以绍兴市为例_百度学术
                    </a>

            </h3>


<div class="op-xueshu-detail-subtitle c-gray"><a href="http://www.baidu.com/link?url=D6ukhURUB3dVbDtt8u6f_8Z9NCIuZskWX8eLruTJcjbTfFJJH0L-LiZSfLkwStBd-CPJxeboEuAQLhrDokbDHmGSLj8MfxtjyJIerAMHFNmCPCiXvJe8zVC4k9Di3fZO8oKjaN7oIbLwPnNa1Y1PLa4Wt_OMTM6gCOC5s6M7VSZWJh49FqxePifv3GpB2oPne8ACA8chG9AOgnDiLRsC1_" target="_blank" class="c-gray">戴大新</a>&nbsp;,&nbsp;<a href="http://www.baidu.com/link?url=D6ukhURUB3dVbDtt8u6f_8Z9NCIuZskWX8eLruTJcjbTfFJJH0L-LiZSfLkwStBd-CPJxeboEuAQLhrDokbDHh255jsZ8bDVFb9j1NBuUUXT1ycXODHpX_THEYf9b2-DFnAHH84lG0qJoRVNk30FNNsw9BYAOVMjLzvdRq5IVsdwU3uixrTgFUYMGLa_aUKUXKHkbHMKM0KS0dER37hGcK" target="_blank" class="c-gray">魏建慧</a>&nbsp;-&nbsp;江南论坛&nbsp;-&nbsp;2019</div><p><em>市域社会治理</em>,是指以地级市的行政区域为范围,依靠党 
委政府,社会组织,企事业单位及个人等主体,创新社会治理机制,对辖区内的人民,事务,组织等进行管理和服务的总和或过...</p><div ><a class="c-showurl OP_LOG_BTN" data-click="{rsv_click_type:'xueshu_detail_pc_showurl'}" href="http://www.baidu.com/link?url=D6ukhURUB3dVbDtt8u6f_8Z9NCIuZskWX8eLruTJcjbTfFJJH0L-LiZSfLkwStBdAv_CWEUt9BsSKlpc_jI9W5PoLTMceKhrZBLcbS6rPga5s_xGKyy1PNte9dUyBVW_42HLaEBuH-yNC_A5G4HAjtMoMQaAr5v3XEUqeiZ6dWrNMQCCuwYkyhbX1w2Hfg83EK9a9vgt0DjbtYiTpNWcuP1PalgVRGHstKFw4J2FxG_6W5Cebg06LAZRB6eFAgV8heQmndVIhrqBr8CsWRYMScZQXXhx6yq46qDh2DWjHajZExtzEo3l67VCAugrsPQlixcH5QfcCjpEyQ1eb7bBZjF5mTf8nkOZbdziLWBZhBvtVoCFOlDk2GbOZPNYm7WVcZOLqTQl_Sh3RI6BMdy_po__0s89EI0uMjZJt81Ri37SE8pZ_t0ryeuKsLoBT5R-uwuCzbYEAfSngynpRs7ctA5UD_KOlcbp4cnAo_nbjZe3WM6SwL0_nZMNDrMW9TMm" target="_blank" style="text-decoration: none;"><span class="op-xueshu-detail-wrap"><img class="op-xueshu-detail-icon c-gap-right-small" src="https://timg01.bdimg.com/timg?pacompress=&imgtype=3&sec=1439619614&di=45267d415b6ec33d25f50cacf6ffc436&quality=90&size=b870_10000&src=http%3A%2F%2Fpic.rmb.bdstatic.com%2F016a609f8e4ce54572b9e06f53c59274.png">百度学术</span></a><span class="c-tools" id="tools_10174712262008669711_12" data-tools="{title:'市域社会治理现代化路径研究——以绍兴市为例_百度学术',url:'http://xueshu.baidu.com/s?wd=paperu
ri:(1v2r00k0pt260tj0ss2w0pa0wy042721)&amp;filter=sc_long_sign&amp;sc_ks_para=q%3D%E5%B8%82%E5%9F%9F%E7%A4%BE%E4%BC%9A%E6%B2%BB%E7%90%86%E7%8E%B0%E4%BB%A3%E5%8C%96%E8%B7%AF%E5%BE%84%E7%A0%94%E7%A9%B6%E2%80%94%E2%80%94%E4%BB%A5%E7%BB%8D%E5%85%B4%E5%B8%82%E4%B8%BA%E4%BE%8B&amp;tn=SE_baiduxueshu_c1gjeupa&amp;ie=utf-8&amp;sc_us=12010585807122850015'}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></span></div>

</div>



















                <div class="result c-container " id="13" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DD6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'ECFDF7FF'
                                                                                                }"
        href = "http://www.baidu.com/link?url=FCq63Pv5w4FX_OVyOXrekqpae_aCcI1ZcCL5gs4enoS3MRRE_1N90D9GNoCd_fSKOFfR-7cgsnA6zqxS_-jtIJWL2hJIi1qTDFpXrFvofDzMkyfdoshSGUm36cY9OGHOrE2xL7-x12537jSr08JzZq"

                            target="_blank"

                >达州市创新“3485”模式 加快推进<em>市域社会治理现代化</em>- 四川省人民...</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年5月6日&nbsp;-&nbsp;</span>达州市认真贯彻落实党
的十九届四中全会精神,对照中央和省委、市委部署要求,创新“3485”模式,推动社会治理和服务重心下移、资源下沉,提升<em>市域社会治理</em>...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=FCq63Pv5w4FX_OVyOXrekqpae_aCcI1ZcCL5gs4enoS3MRRE_1N90D9GNoCd_fSKOFfR-7cgsnA6zqxS_-jtIJWL2hJIi1qTDFpXrFvofDzMkyfdoshSGUm36cY9OGHOrE2xL7-x12537jSr08JzZq" class="c-showurl " style="text-decoration:none;position:relative;">四川省人民政府</span></a><div class="c-tools " id="tools_9460014566025870833_13" data-tools='{"title":"达州市创新“3485”模式 加快推进市域社会治理现代化- 四川省人民...","url":"http://www.baidu.com/link?url=FCq63Pv5w4FX_OVyOXrekqpae_aCcI1ZcCL5gs4enoS3MRRE_1N90D9GNoCd_fSKOFfR-7cgsnA6zqxS_-jtIJWL2hJIi1qTDFpXrFvofDzMkyfdoshSGUm36cY9OGHOrE2xL7-x12537jSr08JzZq"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9f65cb4a8c8507ed4fece763105392230e54f7226dc0d0622a89d75f93130a1d5a65e1bd23220d07d1c378611cac4b5aeef72b7135042bb086c88a4adfa695352b95763f24188d1119d34fffcb4223c773c050bda54ee1b8fb30c3ffd1d4d95302cb44050dc1aacd055e008f32b0423ef4d7ea5f635d07bb9d2713fe4e0329882230a136f2f7456b10f186ca2b4fd42ba0766793b844c72963b704d3690c2534b73dc51f2722279049308b372a05e2fc2d973d0a34&p=906dc64ad49f11a05bed913b514989&newp=882a9645d39714fd0ab6c7710f0c80231610db2151d4d2116b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c57479dc72662247&s=98dce83da57b0395&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=13"
                        target="_blank"
                    class="m ">百度快照</a></div></div>


















                <div class="result c-container " id="14" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'67BF7DDF'
                                                                                                }"
        href = "http://www.baidu.com/link?url=Z0nnrBbX0j8NiGQZV7WGp_K_F7QEOTCgf5CwnzzufvrByjkH20kJ4vyWDt9AqbvPNvDX4Da_m4MsUa8k4lecVYOsdWvI50exlCDxx22N9JW"

                            target="_blank"

                >张掖<em>市域社会治理现代化</em>试点工作取得显著成效-张掖-每日甘肃网</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月3日&nbsp;-&nbsp;</span>每日甘肃网6月3日讯 据张 
掖日报报道(记者杨静文)近日,在全省<em>市域社会治理现代化</em>试点工作推进会议上,张掖市作为全省第一批<em>市域 
社会治理现代化</em>试点城市,...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=Z0nnrBbX0j8NiGQZV7WGp_K_F7QEOTCgf5CwnzzufvrByjkH20kJ4vyWDt9AqbvPNvDX4Da_m4MsUa8k4lecVYOsdWvI50exlCDxx22N9JW" class="c-showurl " style="text-decoration:none;position:relative;">zy.gansudaily.com.cn/system/20...</a><div class="c-tools " id="tools_15749676355598894866_14" data-tools='{"title":"张掖市域社会治理 
现代化试点工作取得显著成效-张掖-每日甘肃网","url":"http://www.baidu.com/link?url=Z0nnrBbX0j8NiGQZV7WGp_K_F7QEOTCgf5CwnzzufvrByjkH20kJ4vyWDt9AqbvPNvDX4Da_m4MsUa8k4lecVYOsdWvI50exlCDxx22N9JW"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b87984aea5868c08128d055621debb9c5a77459730e71446b2fa84095e1e56edb22d65fd5975349d275db446e4e2316b0187fcda5647d42aa7204bd1f06b&p=806acf1eba934eac58ebd72d021489&newp=8c7cd215d9c041a902aed72d021486231610db2151d3d601298ffe0cc4241a1a1a3aecbf2d251000d4c3786d03a84a56e1f734783d0034f1f689df08d2ecce7e75cb64&s=b6d767d2f8ed5d21&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=14"
                        target="_blank"
                    class="m ">百度快照</a></div></div>


















                <div class="result c-container " id="15" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DD6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'7F7AF5F9'
                                                                                                }"
        href = "http://www.baidu.com/link?url=-gtwxmS-pjBMgNj4UL2KRqZTDIT0IFgn_03WxrxbFRLMtdeOFZ3tkij_KokPs62662ED4IM76O_J_2MhrwDNB7Zg3FlreRY0KqzZ2lTebaS"

                            target="_blank"

                >江西省人民政府 江西要闻 加快推进<em>市域社会治理现代化</em></a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月5日&nbsp;-&nbsp;</span>6月3日至4日,省委常委、省委政法委
书记尹建业在赣州市南康区、大余县专题调研<em>市域社会治理现代化</em>工作。他强调,要以列入全国首批<em>市域社会
治理现代化</em>试...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=-gtwxmS-pjBMgNj4UL2KRqZTDIT0IFgn_03WxrxbFRLMtdeOFZ3tkij_KokPs62662ED4IM76O_J_2MhrwDNB7Zg3FlreRY0KqzZ2lTebaS" class="c-showurl " style="text-decoration:none;position:relative;">www.jiangxi.gov.cn/art/2020/6/...</a><div class="c-tools " id="tools_10103689104747499946_15" data-tools='{"title":"江西省人民政府 江西要闻 加快
推进市域社会治理现代化","url":"http://www.baidu.com/link?url=-gtwxmS-pjBMgNj4UL2KRqZTDIT0IFgn_03WxrxbFRLMtdeOFZ3tkij_KokPs62662ED4IM76O_J_2MhrwDNB7Zg3FlreRY0KqzZ2lTebaS"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9f65cb4a8c8507ed4fece763105392230e54f73b67848c40358f8448e4310605506694ea7b3f434495d87c6501ad5458f7f72b21774537b0efc9834bcabce62b258c2331751f914165895ff09552609c60c655fede6af0ccf22592dec5a5de4327ca44040a9781804d7711dd1f800346e0b19838022914ad9c47728e5b605ee93441c65088942518039686db4b38b03dd11106e7df22c34a&p=8b2a9715d9c041a944aef82752528b&newp=9e6dcd1f85cc43e40cbd9b7d0d108c231610db2151d4db176b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c57479c373&s=c81e728d9d4c2f63&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=15"
                        target="_blank"
                    class="m ">百度快照</a></div></div>


















                <div class="result c-container " id="16" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DD6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'BDFEBBEB'
                                                                                                }"
        href = "http://www.baidu.com/link?url=Ju28ukVuL4KfPatwxZ4fv0q0H3yOvXgvQ8JAUWYWyDKZLfWp9CY5jgfJB3ThJbb3d8naUhJK2t-Ld1GaHLUvza"

                            target="_blank"

                ><em>市域社会治理现代化</em>,这个新词儿到底是啥意思?_体系</a></h3><div class="c-row c-gap-top-small"><div class="general_image_pic c-span6" ><a class="c-img6" style="height:75px"
          href="http://www.baidu.com/link?url=Ju28ukVuL4KfPatwxZ4fv0q0H3yOvXgvQ8JAUWYWyDKZLfWp9CY5jgfJB3ThJbb3d8naUhJK2t-Ld1GaHLUvza"
                target="_blank"
      ><img class="c-img c-img6" src="https://dss2.baidu.com/6ONYsjip0QIZ8tyhnq/it/u=4203659606,3748133425&fm=85&app=92&f=JPEG?w=121&h=75&s=45105530794052CC0D7CB5C60100E0B1"  style="height:75px;" /></a></div><div class="c-span18 c-span-last"><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2019年7月3日&nbsp;-&nbsp;</span>2018年6月4日,在延安干部学院新任地市级党委政法委书记培训示范班开班式上,中央政法委秘书长陈一 
新首次正式提出“<em>市域社会治理现代化</em>”的概念,并从治理理念现代化、...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=Ju28ukVuL4KfPatwxZ4fv0q0H3yOvXgvQ8JAUWYWyDKZLfWp9CY5jgfJB3ThJbb3d8naUhJK2t-Ld1GaHLUvza" class="c-showurl " style="text-decoration:none;position:relative;"><style>.nor-src-wrap {position: relative;}
                        .nor-src-icon {vertical-align: middle;width: 14px;height: 14px;border: 1px solid #eee;border-radius: 100%;margin-right: 5px;margin-top: -3px;position: relative;}
                        .nor-src-icon-with-v {margin-right: 10px;}
                        .nor-src-icon-v {display: inline-block;width: 10px;height: 10px;border-radius: 100%;position: absolute;left: 8px;bottom: 0;background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);background-size: 10px 10px;}
                        .nor-src-icon-v.vicon-1 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/red-v.png);}
                        .nor-src-icon-v.vicon-2 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/blue-v.png);}
                        .nor-src-icon-v.vicon-3 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);}</style><span class="nor-src-wrap"><img class="nor-src-icon" src="https://timg01.bdimg.com/timg?pacompress=&amp;imgtype=0&amp;sec=1439619614&amp;autorotate=1&amp;di=45b00e4172635b0c418b72477c38e5c4&amp;quality=90&amp;size=b870_10000&amp;src=http%3A%2F%2Fpic.rmb.bdstatic.com%2Feab014be45108fa8ab9580e993530dd0.png">搜狐网</span></a><div class="c-tools " id="tools_13320425487750752003_16" data-tools='{"title":"市域社会治理现代化,这个新词儿到底是啥意思?_体系","url":"http://www.baidu.com/link?url=Ju28ukVuL4KfPatwxZ4fv0q0H3yOvXgvQ8JAUWYWyDKZLfWp9CY5jgfJB3ThJbb3d8naUhJK2t-Ld1GaHLUvza"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618a8b81894a7ed77d6ebc5d3a8150e9014127af7a1d6051d47d76ef01235e0a29f49174817ceee3267fd5975229822&p=c9769a479f934eac58eb826d154f8c&newp=8b2a975d86cc42af5eb3d5664f649c231610db2151d4d601298ffe0cc4241a1a1a3aecbf2d251000d4c3786d03a84a56e1f734783d0034f1f689df08d2ecce7e37&s=7288251b27c8f0e7&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=16"
                        target="_blank"
                    class="m ">百度快照</a></div></div></div></div>


















                <div class="result c-container " id="17" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'DDCFC6DB'
                                                                                                }"
        href = "http://www.baidu.com/link?url=M_bTXOx5Qt6iQriwC6rOCdTCdhQEoKeTAaUs7TAmaincxKUbcpGT4pv9tIEPzi0FOZwEtvPE7cH0Oow8l00UI_"

                            target="_blank"

                >专题(二)<em>市域社会治理</em>是个啥? - <em>市域社会治理现代化</em>专题 - 中国...</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月4日&nbsp;-&nbsp;</span>党的十九届 
四中全会明确提出“加快推进<em>市域社会治理现代化</em>”。市域社会治理是国家治理的重要维度,在国家治理体系中具有
承上启下的枢纽作用。 市域社会...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=M_bTXOx5Qt6iQriwC6rOCdTCdhQEoKeTAaUs7TAmaincxKUbcpGT4pv9tIEPzi0FOZwEtvPE7cH0Oow8l00UI_" class="c-showurl " style="text-decoration:none;position:relative;">www.csgpc.org/bencan...php?fid...</a><div class="c-tools " id="tools_10237569640625103109_17" data-tools='{"title":"专题(二)市域社会治理是个啥? - 市域社 
会治理现代化专题 - 中国...","url":"http://www.baidu.com/link?url=M_bTXOx5Qt6iQriwC6rOCdTCdhQEoKeTAaUs7TAmaincxKUbcpGT4pv9tIEPzi0FOZwEtvPE7cH0Oow8l00UI_"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b577d6b9d2a48e15068802127af7adcb0f1d449d33a1476db5ed8e48621d4af8fa3111ab02213e9e5543b744fbf7326a5cd7e1dc2a48c12cdb&p=ce67c64ad4db11a05bec933a4e08c6&newp=91759a46d7c457b40cbb882d021482231610db2151d6d11f6b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c574799e3c602600&s=cfcd208495d565ef&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=17"
                        target="_blank"
                    class="m ">百度快照</a></div></div>


















                <div class="result c-container " id="18" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'B7FFFCE7'
                                                                                                }"
        href = "http://www.baidu.com/link?url=Oe7q8KYG12-NX1q6ZC7K7HHs24tQ4mBhAZ0c5IsE5CjOC9Bkcag86ywWJhCESvCJSixPtxVmwg7CeSxnvtjhRK"

                            target="_blank"

                >推进<em>市域治理现代化</em> 建设更高水平平安合肥_合肥市人民政府</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月20日&nbsp;-&nbsp;</span>6月18日,合肥市创建全国<em> 
市域社会治理现代化</em>试点城市动员部署会在市政务中心召开,市委副书记郭强出席会议并讲话。他强调,要深入学习贯 
彻习近平总书记关...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=Oe7q8KYG12-NX1q6ZC7K7HHs24tQ4mBhAZ0c5IsE5CjOC9Bkcag86ywWJhCESvCJSixPtxVmwg7CeSxnvtjhRK" class="c-showurl 
" style="text-decoration:none;position:relative;">www.hefei.gov.cn/ssxw/zwyw/105...</a><div class="c-tools " id="tools_4360335013975936206_18" data-tools='{"title":"推进市域治理现代化 建设更高水平平安合肥_合肥市人民 
政府","url":"http://www.baidu.com/link?url=Oe7q8KYG12-NX1q6ZC7K7HHs24tQ4mBhAZ0c5IsE5CjOC9Bkcag86ywWJhCESvCJSixPtxVmwg7CeSxnvtjhRK"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span 
class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b577d6b9d2a48503079d08127af7a5d61e1714bd3eac0970a2ecdc555d1a5bfff03367f85b7122992641a1478eba743159&p=81759a46d7c016f80abe9b7c475995&newp=8a769a47a48b14f708e2977f0e4f8b231610db2151d7d6136b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c57479fa6f6128&s=e3f3064ac424a80e&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=18"
                        target="_blank"
                    class="m ">百度快照</a></div></div>


















                <div class="result c-container " id="19" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'CB5FB9FF'
                                                                                                }"
        href = "http://www.baidu.com/link?url=R9vcdsmEAncjJZm0Q5_pVKeE5LDgSdKCq6wCODx4PhZZEoSp1e5TI024jRJRnsxGKg14SrJ3Y84SbwlNWNu1eV8fMXR9W3WZt_gGzwlKYq3"

                            target="_blank"

                >全市<em>市域社会治理现代化</em>工作现场推进会召开</a></h3><div class="c-row c-gap-top-small"><div class="general_image_pic c-span6" ><a class="c-img6" style="height:75px"
          href="http://www.baidu.com/link?url=R9vcdsmEAncjJZm0Q5_pVKeE5LDgSdKCq6wCODx4PhZZEoSp1e5TI024jRJRnsxGKg14SrJ3Y84SbwlNWNu1eV8fMXR9W3WZt_gGzwlKYq3"
                target="_blank"
      ><img class="c-img c-img6" src="https://dss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=2189432748,564677081&fm=173&app=49&f=JPEG?w=312&h=208&s=EF12658D728B30E36A090DC6030010B3"  style="height:75px;" /></a></div><div class="c-span18 c-span-last"><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月18日&nbsp;-&nbsp;</span>6月17日,全市<em>市域社会治理现代化</em>工作现场推进会在文登区召开,传达全省<em>市域社会治 
理现代化</em>试点工作视频会议精神和省委常委、政法委书记林峰海同志在我市调研时的指示...</div><div class="f13 
 se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=R9vcdsmEAncjJZm0Q5_pVKeE5LDgSdKCq6wCODx4PhZZEoSp1e5TI024jRJRnsxGKg14SrJ3Y84SbwlNWNu1eV8fMXR9W3WZt_gGzwlKYq3" class="c-showurl " style="text-decoration:none;position:relative;"><style>.nor-src-wrap {position: relative;}
                        .nor-src-icon {vertical-align: middle;width: 14px;height: 14px;border: 1px solid #eee;border-radius: 100%;margin-right: 5px;margin-top: -3px;position: relative;}
                        .nor-src-icon-with-v {margin-right: 10px;}
                        .nor-src-icon-v {display: inline-block;width: 10px;height: 10px;border-radius: 100%;position: absolute;left: 8px;bottom: 0;background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);background-size: 10px 10px;}
                        .nor-src-icon-v.vicon-1 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/red-v.png);}
                        .nor-src-icon-v.vicon-2 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/blue-v.png);}
                        .nor-src-icon-v.vicon-3 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);}</style><span class="nor-src-wrap"><img class="nor-src-icon nor-src-icon-with-v" src="https://cambrian-images.cdn.bcebos.com/7f7a8d7b247d3aa430010f10a5765239_1630526374013150.jpeg@w_100,h_100"><span class="nor-src-icon-v vicon-2"></span>海报新闻</span></a><div class="c-tools " id="tools_10826130747089631106_19" data-tools='{"title":"全市市域社会治理现代化工作现场推进会召开","url":"http://www.baidu.com/link?url=R9vcdsmEAncjJZm0Q5_pVKeE5LDgSdKCq6wCODx4PhZZEoSp1e5TI024jRJRnsxGKg14SrJ3Y84SbwlNWNu1eV8fMXR9W3WZt_gGzwlKYq3"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618a8b81894a7ed62c0f58a888c0e009744050dd0a3d00c4703ca18a1496efee78e49610446adec4666fb5d7c2c94284bb240fee1306b008bf3da595bc32e95355180ae43f17c49e344e4081b2010fd0ca609276461&p=c362c415d9c043ae1fb9cd2d02149d&newp=9c769a47929209f601bd9b7d0f178e231610db2151d4d3136b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c57479cc767c2945&s=1ff1de774005f8da&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=19"
                        target="_blank"
                    class="m ">百度快照</a></div></div></div></div>


















                <div class="result c-container " id="20" srcid="1599" tpl="se_com_default"  data-click="{'rsv_bdr':'0' }"  ><h3 class="t" ><a
                data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595927670',
                                                'y':'7E25DDDB'
                                                                                                }"
        href = "http://www.baidu.com/link?url=-JL9W4zrN6duV1x34txtuyGJ2qY5viZS4HCdIz1dpPBAbFBltUeM_5NvYTmuEF8VVw40c_Uy2mmtFKBZL5dfKfygs7MYMYtK_LZv03WdK-8gqLgOTy6fO2EDSB7qMilp"

                            target="_blank"

                >金华市政府- <em>市域社会治理现代化</em>如何稳步推进</a></h3><div class="c-row c-gap-top-small"><div class="general_image_pic c-span6" ><a class="c-img6" style="height:75px"
          href="http://www.baidu.com/link?url=-JL9W4zrN6duV1x34txtuyGJ2qY5viZS4HCdIz1dpPBAbFBltUeM_5NvYTmuEF8VVw40c_Uy2mmtFKBZL5dfKfygs7MYMYtK_LZv03WdK-8gqLgOTy6fO2EDSB7qMilp"
                target="_blank"
      ><img class="c-img c-img6" src="https://dss1.bdstatic.com/6OF1bjeh1BF3odCf/it/u=1980592187,880886268&fm=85&app=92&f=JPEG?w=121&h=75&s=F99161954C5054DAC80CF5CE0300F071"  style="height:75px;" /></a></div><div class="c-span18 c-span-last"><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年3月30日&nbsp;-&nbsp;</span>习近平总书记指出,“这次抗击新冠肺炎疫情,是对国家治理体系和治理能力的一次大考”。<em>市域社会
治理</em>是国家治理的重要基础,是国家治理体系与治理能力<em>现代化</em>承上启下的...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=-JL9W4zrN6duV1x34txtuyGJ2qY5viZS4HCdIz1dpPBAbFBltUeM_5NvYTmuEF8VVw40c_Uy2mmtFKBZL5dfKfygs7MYMYtK_LZv03WdK-8gqLgOTy6fO2EDSB7qMilp" class="c-showurl " style="text-decoration:none;position:relative;">www.jinhua.gov.cn/113307000025...</a><div class="c-tools " id="tools_18005358456910554366_20" data-tools='{"title":"金华市政府- 市域社会治理现代化如何稳步推进 ","url":"http://www.baidu.com/link?url=-JL9W4zrN6duV1x34txtuyGJ2qY5viZS4HCdIz1dpPBAbFBltUeM_5NvYTmuEF8VVw40c_Uy2mmtFKBZL5dfKfygs7MYMYtK_LZv03WdK-8gqLgOTy6fO2EDSB7qMilp"}'><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a 
data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b577d6b9d2a4870f0f9014566d8087de074403ca18a1482ce0a59849175a12b8ef3265f852772e942834ab11b2a56a730783f6df5e4dde6cd0631195ad34b13c05b267a2181d2345aa4fa60f252527e0387db91b63&p=cb769a4786cc41ae0fa2c63e4a64&newp=8b2a975bc7db15ff57ed95341c4a92695d0fc20e3ad5d301298ffe0cc4241a1a1a3aecbf2d251000d4c3786d03a84a56e1f734783d0034f1f689df08d2ecce7e31993f60&s=c4ca4238a0b92382&user=baidu&fm=sc&query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&qid=93b6eb870000003b&p1=20"
                        target="_blank"
					class="m ">百度快照</a></div></div></div></div>`

// var re_baidu = `<div class="result c-container "[\d\D]*?href = "([^"]*?)"[\d\D]*?>([\d\D]*?)</a></h3>[\d\D]*?<div class="c-abstract"><span class=" newTimeFactor_before_abs  m">([^<]*?)&nbsp;-&nbsp;</span>([\d\D]*?)</div>`

var body2 = `class="result c-container " id="4" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:4}"><h3 class="t"><a data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'3CFF46B7'
                                                                                                }" href="http://www.baidu.com/link?url=-i-lJNoAdS6Kudsx2F85j50UReOfwI2dRoJkVMBC7dp0BTw_MPQ5rH6sUEZXShwh8BbJMh_4fPmu-FcSDCso_gwgZUGB77jtNmULpBbfDKa" target="_blank">关于推进<em>市域社会治理现代化</em>的实施意见 出台 - 河北新闻网</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月5日&nbsp;-&nbsp;</span>日前,承德市出台《关于推进<em>市域社会治理现代化</em>的实施意见》,将以
开展全国<em>市域社会治理现代化</em>试点为抓手,积极探索具有时代特征、承德特色的市域社会治...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=-i-lJNoAdS6Kudsx2F85j50UReOfwI2dRoJkVMBC7dp0BTw_MPQ5rH6sUEZXShwh8BbJMh_4fPmu-FcSDCso_gwgZUGB77jtNmULpBbfDKa" class="c-showurl " style="text-decoration:none;position:relative;"><style>.nor-src-wrap {position: relative;}
                        .nor-src-icon {vertical-align: middle;width: 14px;height: 14px;border: 1px solid #eee;border-radius: 
100%;margin-right: 5px;margin-top: -3px;position: relative;}
                        .nor-src-icon-with-v {margin-right: 10px;}
                        .nor-src-icon-v {display: inline-block;width: 10px;height: 10px;border-radius: 100%;position: absolute;left: 8px;bottom: 0;background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);background-size: 10px 10px;}
                        .nor-src-icon-v.vicon-1 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/red-v.png);}
                        .nor-src-icon-v.vicon-2 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/blue-v.png);}
                        .nor-src-icon-v.vicon-3 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);}</style><span class="nor-src-wrap"><img class="nor-src-icon nor-src-icon-with-v" src="https://cambrian-images.cdn.bcebos.com/ff3065066fff541a23d1698e0d0a85a1_1571964251298255.jpeg@w_100,h_100"><span class="nor-src-icon-v vicon-2"></span>河北新闻网</span></a><div class="c-tools " id="tools_16861998499912839986_4" data-tools="{&quot;title&quot;:&quot;关于推进市 
域社会治理现代化的实施意见 出台 - 河北新闻网&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=-i-lJNoAdS6Kudsx2F85j50UReOfwI2dRoJkVMBC7dp0BTw_MPQ5rH6sUEZXShwh8BbJMh_4fPmu-FcSDCso_gwgZUGB77jtNmULpBbfDKa&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618a8b81894a7ed62c0f58a888c0e009744050dd0a3d00c4703ca18a1496efee78e49610446adec4666fb5d7d2d982541b747f2e63168038af5dc585bc32e95355180ae43f17c49e344e4081b2010fd0ca609276461&amp;p=9b3fc64ad4d015b708e2917f4e55&amp;newp=8e6acc1487d512a05abd9b7b0c6492695d0fc20e38d4d301298ffe0cc4241a1a1a3aecbf2d251000d4c3786d03a84a56e1f734783d0034f1f689df08d2ecce7e&amp;s=cfcd208495d565ef&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=4" target="_blank" class="m ">百度快照</a></div></div>
















                                                                                                                        <div 
class="result c-container " id="5" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:5}"><h3 class="t"><a data-click="{
                        'F':'778F37EA',
                        'F1':'9D63F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'7FFFFF3E'
                                                                                                }" href="http://www.baidu.com/link?url=-TbFlQme9isK9kDy--1Mc4FLfEYz4UEBumXNajFEKJnB2EnBSIRkLAFeC0eEL7Sd_ujfrMjhmaljptbeJdYBoJ4aW-k9MHq0Pq1kgcatF0u" target="_blank">全面安排部署<em>市域社会治理现代化</em>试点工作-银川日报·轻松阅读</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">1天内&nbsp;-&nbsp;</span>本报讯(记者范晓儒)7月27日,我市召开<em>市域社会治理现代化</em>试点工作领导 
小组会议,传达学习全国<em>市域社会治理现代化</em>工作会议、平安宁夏建设协调小组第一次...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=-TbFlQme9isK9kDy--1Mc4FLfEYz4UEBumXNajFEKJnB2EnBSIRkLAFeC0eEL7Sd_ujfrMjhmaljptbeJdYBoJ4aW-k9MHq0Pq1kgcatF0u" class="c-showurl " style="text-decoration:none;position:relative;">szb.ycen.com.cn/epaper/ycrb/ht...</a><div class="c-tools " id="tools_12547437740466523332_5" data-tools="{&quot;title&quot;:&quot;全面安排部署市 
域社会治理现代化试点工作-银川日报·轻松阅读&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=-TbFlQme9isK9kDy--1Mc4FLfEYz4UEBumXNajFEKJnB2EnBSIRkLAFeC0eEL7Sd_ujfrMjhmaljptbeJdYBoJ4aW-k9MHq0Pq1kgcatF0u&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9f65cb4a8c8507ed4fece76310579f360e54f7286d808c027fa3c215cc735b36163afeec6571525393d8373641ff5406acaf686f370120b58cc8fe48d8a6922232d97a68364bda0705d46dafc04d2fc137902db3e946f3ffad72c5a1c5a2ac4325c844040a97868a4d7414dd6e800340e8b1ee4d022f60ad9a3372fe29605f9b3431c15088e225197196f7ad4b3db73da66506e7a922c44d05b463b36f6b3337d4&amp;p=8b2a971d998b0bdd0cbd9b7d0f4f&amp;newp=9a6ece0d85cc43e742bd9b7d0f0592695c02dc3051d6d001298ffe0cc4241a1a1a3aecbb24241502d9c478610abb0f31aba7747d605f76ff81&amp;s=e46bc064f8e92ac2&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=5" target="_blank" class="m ">百度快照</a></div></div>
















                                                                                                                        <div 
class="result c-container " id="6" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:6}"><h3 class="t"><a data-click="{
                        'F':'778F37EA',
                        'F1':'9D53F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'BFE7FDED'
                                                                                                }" href="http://www.baidu.com/link?url=fJ5P3mDdJxPJPBMMaMXI6-MCCJ9WznH5DcGFgS6ZeNXQgpMk6gg8NQufOc07M9yiF83yCmEiX0wHNKpLro9mVETpb44T1w1WWf5MWk-FKXtIWcv-H0ZP_Aodt-4HxoQp" target="_blank">打造<em>市域社会治理现代化</em>先行示范区 _中国经济网——国家经济门户</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">1天前&nbsp;-&nbsp;</span><em>市域社会治理现代化</em>是在城市范围开展的地方治
理创新活动,是国家治理现代化的目标及要求在市域的体现。广州<em>市域社会治理现代化</em>是发挥市级统筹协调资源...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=fJ5P3mDdJxPJPBMMaMXI6-MCCJ9WznH5DcGFgS6ZeNXQgpMk6gg8NQufOc07M9yiF83yCmEiX0wHNKpLro9mVETpb44T1w1WWf5MWk-FKXtIWcv-H0ZP_Aodt-4HxoQp" class="c-showurl " style="text-decoration:none;position:relative;">views.ce.cn/view/ent/202007/27...</a><div class="c-tools " id="tools_13351190852463847271_6" data-tools="{&quot;title&quot;:&quot;打造市域社会治理现代化先行示范区 _中国经济网——国家经济门户&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=fJ5P3mDdJxPJPBMMaMXI6-MCCJ9WznH5DcGFgS6ZeNXQgpMk6gg8NQufOc07M9yiF83yCmEiX0wHNKpLro9mVETpb44T1w1WWf5MWk-FKXtIWcv-H0ZP_Aodt-4HxoQp&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b469c4eb93c4df23029d44050dd1ac961e5b438f72a74877fea69b48175d15a7ed3578b95975299d2145b642eee7466f0087f4df5f48c53dd01650cde96aee&amp;p=9b3fc64ad4d015b708e2947b4e55&amp;newp=8e6acc1487d512a05abd9b7e086492695912c10e3fd5d301298ffe0cc4241a1a1a3aecbb24261a02d4c27f630abb0f31aba7747d605f76&amp;s=cfcd208495d565ef&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=6" 
target="_blank" class="m ">百度快照</a></div></div>
















                                                                                                                        <div 
class="result c-container " id="7" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:7}"><h3 class="t"><a data-click="{
                        'F':'778717EA',
                        'F1':'9D53F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'FBFFDDCE'
                                                                                                }" href="http://www.baidu.com/link?url=byEWfYiYZnoec1SZSy-E8Gc1N-gca87HvjVxIKZFC4Zzhcm4fLeLULDk-7POhRZCLAgsY_zTxh1Sr2T4RCr-VK6Gt1dP2oJBT7zO9HTqcIq" target="_blank">为推进<em>市域社会治理现代化</em>贡献检察力量_中华人民共和国最高人民检...</a></h3><div class="c-row c-gap-top-small"><div class="general_image_pic c-span6"><a class="c-img6" style="height:75px" href="http://www.baidu.com/link?url=byEWfYiYZnoec1SZSy-E8Gc1N-gca87HvjVxIKZFC4Zzhcm4fLeLULDk-7POhRZCLAgsY_zTxh1Sr2T4RCr-VK6Gt1dP2oJBT7zO9HTqcIq" target="_blank"><img class="c-img c-img6" src="https://dss2.bdstatic.com/6Ot1bjeh1BF3odCf/it/u=536729511,2390134936&amp;fm=85&amp;app=2&amp;f=JPEG?w=121&amp;h=75&amp;s=8EC58346EC2910155B3835AA0300E01C" style="height:75px;"></a></div><div class="c-span18 c-span-last"><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2天前&nbsp;-&nbsp;</span>而推进社会治理现代化,法治保障是重要环节和 
重要内容,检察机关应充分发挥“四大检察”职能,以高质量的法律监督工作推动<em>市域社会治理现代化</em>,贡献检察智慧。 市域社会...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=byEWfYiYZnoec1SZSy-E8Gc1N-gca87HvjVxIKZFC4Zzhcm4fLeLULDk-7POhRZCLAgsY_zTxh1Sr2T4RCr-VK6Gt1dP2oJBT7zO9HTqcIq" class="c-showurl " style="text-decoration:none;position:relative;">中华人民共和国最高人...</a><div class="c-tools " id="tools_5379526486927500726_7" data-tools="{&quot;title&quot;:&quot;为推进市域社会治理现代化贡献检察力量_中华人民共和国最高人民检察院&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=byEWfYiYZnoec1SZSy-E8Gc1N-gca87HvjVxIKZFC4Zzhcm4fLeLULDk-7POhRZCLAgsY_zTxh1Sr2T4RCr-VK6Gt1dP2oJBT7zO9HTqcIq&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618a8b81894a7ed77d6ebc5d3a815118844050dd5adcf4d00639b33ed5573a1bbc7165e070dbaef3067fd5c6a6f9f2140b445fce0367900f5f0d85d47c12cc76166d6f573ef60&amp;p=9d3bc54ad7c51df90da6c7710f43&amp;newp=9f3bc54ad7c513fe0be2966f174292695912c10e3fd5da01298ffe0cc4241a1a1a3aecbb24261a02d4c27f630abb0f31aba7747d605f76e5ca&amp;s=cfcd208495d565ef&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=7" target="_blank" class="m ">百度快照</a></div></div></div></div>
















                                                                                                                        <div 
class="result c-container " id="8" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:8}"><h3 class="t"><a data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'F626EBBB'
                                                                                                }" href="http://www.baidu.com/link?url=3L0zEoSK6j5e5GsRc-rQZJ-GBDW172_XUXzznufNjiU0CvSMktGQBmC-TOPq8MwwO0OCDVR6MaGQuhUZjhYtC9kSrQI_WPgsJdUEH1dheH9Ybwc_R-OJGWFc5eL7UIns" target="_blank">构建城市基层党建新格局 推进<em>市域社会治理现代化</em>--党建-人民网</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月4日&nbsp;-&nbsp;</span>党的十九届四中全会强调,要构建基层社会治理新 
格局,加快推进<em>市域社会治理现代化</em>。这些重要指示和一系列重大部署,为我们加强和改进城市基层党建工作提...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=3L0zEoSK6j5e5GsRc-rQZJ-GBDW172_XUXzznufNjiU0CvSMktGQBmC-TOPq8MwwO0OCDVR6MaGQuhUZjhYtC9kSrQI_WPgsJdUEH1dheH9Ybwc_R-OJGWFc5eL7UIns" class="c-showurl " style="text-decoration:none;position:relative;">dangjian.people.com.cn/GB/n1/2...</a><div class="c-tools " id="tools_10684658238191265566_8" data-tools="{&quot;title&quot;:&quot;构建城市基层党建新格局 推进市域社会治理现代化--党建-人民网&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=3L0zEoSK6j5e5GsRc-rQZJ-GBDW172_XUXzznufNjiU0CvSMktGQBmC-TOPq8MwwO0OCDVR6MaGQuhUZjhYtC9kSrQI_WPgsJdUEH1dheH9Ybwc_R-OJGWFc5eL7UIns&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7a661cffb8a888c0844ca24472dddb2d50d1714bd3ead4b26e3d1c814082a60a7b13378ff5b772b822144b441e4b1316d0283fddd4b4cb52bd3641090a836b12912c249e240453345d417f551462144a11c30fe323e0692ea5eec2f2a&amp;p=8761c70396934eab53bcc7710f5f&amp;newp=87759a41dd9a06f50be2966f1e4492695d0fc20e3bd7d501298ffe0cc4241a1a1a3aecbf2d251000d4c3786d03a84a56e1f734783d0034f1f689df08d2ecce7e76cf7d&amp;s=cfcd208495d565ef&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=8" target="_blank" class="m ">百度快照</a></div></div>
















                                                                                                                        <div 
class="result c-container " id="9" srcid="1599" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;p5&quot;:9}"><h3 class="t"><a data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'FDDF36E7'
                                                                                                }" href="http://www.baidu.com/link?url=jc1t-5ppmwLbuo1gmX2dB36KIvgq9oS4V7EJlJfXlxOuWeECLp-tS-knILNBKOBtHi_3pD7EGpBzebf0AVg50u308NYl8Eu8xKb2N4ufX43" target="_blank"><em>市域社会治理现代化</em>的五个维度_中国网</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2019年12月12日&nbsp;-&nbsp;</span>社会治理现代化是国家治理体系和治理能力现代化的重要组成部分。党的十九届四中全会特别指出
,要加快推进<em>市域社会治理现代化</em>。推进<em>市域社会治理现代化</em>,...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=jc1t-5ppmwLbuo1gmX2dB36KIvgq9oS4V7EJlJfXlxOuWeECLp-tS-knILNBKOBtHi_3pD7EGpBzebf0AVg50u308NYl8Eu8xKb2N4ufX43" class="c-showurl " style="text-decoration:none;position:relative;"><style>.nor-src-wrap {position: 
relative;}
                        .nor-src-icon {vertical-align: middle;width: 14px;height: 14px;border: 1px solid #eee;border-radius: 
100%;margin-right: 5px;margin-top: -3px;position: relative;}
                        .nor-src-icon-with-v {margin-right: 10px;}
                        .nor-src-icon-v {display: inline-block;width: 10px;height: 10px;border-radius: 100%;position: absolute;left: 8px;bottom: 0;background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);background-size: 10px 10px;}
                        .nor-src-icon-v.vicon-1 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/red-v.png);}
                        .nor-src-icon-v.vicon-2 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/blue-v.png);}
                        .nor-src-icon-v.vicon-3 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);}</style><span class="nor-src-wrap"><img class="nor-src-icon" src="https://timg01.bdimg.com/timg?pa=&amp;imgtype=0&amp;sec=1439619614&amp;di=b12a2116212676f80a7a87e2d710b4cd&amp;quality=90&amp;size=b870_10000&amp;src=http%3A%2F%2Fpic.rmb.bdstatic.com%2F1e4a8a6441293e100df0fce6d1ac322a.jpeg">中国网</span></a><div class="c-tools " id="tools_8497806250252311666_9" data-tools="{&quot;title&quot;:&quot;市域社会治理现代化的五个维度_中国网&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=jc1t-5ppmwLbuo1gmX2dB36KIvgq9oS4V7EJlJfXlxOuWeECLp-tS-knILNBKOBtHi_3pD7EGpBzebf0AVg50u308NYl8Eu8xKb2N4ufX43&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618feae6afaa7b577d6b9d2a48e0e089600127af7a1d6051714bd3eac096ca1fdc51348030dfcb76b39a644772b9c2857b631fae02f6d079ca780000a9476967616e3aa32b73c15b015a0081b531ee613&amp;p=cb74cc15d9c040ad34be9b7c565c9e&amp;newp=916ec64ad4dd11a05bec962b130d93231610db2151d7da146b82c825d7331b001c3bbfb423281704d7c27a630bad4e5fe0fb31703d0923a3dda5c91d9fb4c5747996&amp;s=07cdfd23373b17c6&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=9" target="_blank" class="m ">百度快照</a></div></div>
















                                                                                                                        <div 
class="result c-container " id="10" srcid="1525" tpl="se_com_default" data-click="{&quot;rsv_bdr&quot;:&quot;0&quot;,&quot;rsv_cd&quot;:&quot;safe:1|t:1&quot;,&quot;p5&quot;:10}"><h3 class="t"><span class="c-icon c-icon-doc c-gap-icon-right-small"></span><a data-click="{
                        'F':'778317EA',
                        'F1':'9D73F1E4',
                        'F2':'4CA6DE6B',
                        'F3':'54E5263F',
                        'T':'1595937031',
                                                'y':'DB5F7EF9'
                                                                                                }" href="http://www.baidu.com/link?url=aLTOfBGQaJjxu1eGf64Vk8lVbAigMVW4aPZOUKQ9SjAmsIyo6mja7Wc5PjuoTYw72irh92QcdxvTq_rWB33eF4sJyGC9RVogyCSd9eN1_BwEirlVZ8UoTnM91JJvxgm0" target="_blank"><em>市域社会治理现代化</em>典型经验材料_百度文库</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">2020年6月9日&nbsp;-&nbsp;</span>我们将主 动作为,积极探索,努力争创全国<em>市域社会治理现代化</em>示范 城市。 ——整体谋划,以系统思维推进<em>市域社会治理现代化</em>。 围绕 <em>市域社会治理现代化</em>典...</div><div class="f13  se_st_footer"><a target="_blank" href="http://www.baidu.com/link?url=aLTOfBGQaJjxu1eGf64Vk8lVbAigMVW4aPZOUKQ9SjAmsIyo6mja7Wc5PjuoTYw72irh92QcdxvTq_rWB33eF4sJyGC9RVogyCSd9eN1_BwEirlVZ8UoTnM91JJvxgm0" class="c-showurl " style="text-decoration:none;position:relative;"><style>.nor-src-wrap {position: relative;}
                        .nor-src-icon {vertical-align: middle;width: 14px;height: 14px;border: 1px solid #eee;border-radius: 
100%;margin-right: 5px;margin-top: -3px;position: relative;}
                        .nor-src-icon-with-v {margin-right: 10px;}
                        .nor-src-icon-v {display: inline-block;width: 10px;height: 10px;border-radius: 100%;position: absolute;left: 8px;bottom: 0;background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);background-size: 10px 10px;}
                        .nor-src-icon-v.vicon-1 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/red-v.png);}
                        .nor-src-icon-v.vicon-2 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/blue-v.png);}
                        .nor-src-icon-v.vicon-3 {background-image: url(https://b.bdstatic.com/searchbox/icms/searchbox/img/yellow-v.png);}</style><span class="nor-src-wrap"><img class="nor-src-icon" src="https://pic.rmb.bdstatic.com/86207b67cb0c88d82589c0280e5f6682.jpeg">百度文库</span></a><div class="c-tools " id="tools_12894601325846420236_10" data-tools="{&quot;title&quot;:&quot;市域社会治理现代化典型经验材料_百度文库&quot;,&quot;url&quot;:&quot;http://www.baidu.com/link?url=aLTOfBGQaJjxu1eGf64Vk8lVbAigMVW4aPZOUKQ9SjAmsIyo6mja7Wc5PjuoTYw72irh92QcdxvTq_rWB33eF4sJyGC9RVogyCSd9eN1_BwEirlVZ8UoTnM91JJvxgm0&quot;}"><a class="c-tip-icon"><i class="c-icon c-icon-triangle-down-g"></i></a></div><span class="c-icons-outer"><span class="c-icons-inner"><span class="c-trust-as baozhang-new c-icon c-icon-baozhang-new" data_key="1598668362003819882" hint-data="{&quot;label&quot;:&quot;北京百度网讯科技有限公司&quot;,&quot;url&quot;:&quot;https://www.baidu.com/s?wd=%E5%8C%97%E4%BA%AC%E7%99%BE%E5%BA%A6%E7%BD%91%E8%AE%AF%E7%A7%91%E6%8A%80%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8@v&amp;vmp_ec=dcf894083f21e00aca6e03586f146Wh5a1QfX1LdmJab9=Ld=9m1NbaJvmbR31dt27ba63a4lcasX0d0bdf6859b&amp;vmp_ectm=1594959032&amp;from=vs&quot;,&quot;hint&quot;:[], &quot;text&quot;: &quot;\u8be5\u4f01\u4e1a\u5df2\u901a\u8fc7\u5b9e\u540d\u8ba4\u8bc1\uff0c\u67e5\u770b <a href=\&quot;https://www.baidu.com/s?wd=%E5%8C%97%E4%BA%AC%E7%99%BE%E5%BA%A6%E7%BD%91%E8%AE%AF%E7%A7%91%E6%8A%80%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8@v&amp;vmp_ec=dcf894083f21e00aca6e03586f146Wh5a1QfX1LdmJab9=Ld=9m1NbaJvmbR31dt27ba63a4lcasX0d0bdf6859b&amp;vmp_ectm=1594959032&amp;from=vs\&quot; target=\&quot;_blank\&quot;>\u4f01\u4e1a\u6863\u6848<\/a>\u3002<\/br>\u767e\u5ea6\u63a8\u51fa <a href=\&quot;http:\/\/baozhang.baidu.com\/guarantee\/?from=ps\&quot; target=\&quot;_blank\&quot;>\u7f51\u6c11\u6743\u76ca\u4fdd\u969c\u8ba1\u5212<\/a>\uff0c<a href=\&quot;https:\/\/passport.baidu.com\&quot; target=\&quot;_blank\&quot;>\u767b\u5f55<\/a> \u641c\u7d22\u6709\u4fdd\u969c\u3002&quot;}" hint-type="newBao" render="render"></span></span></span>&nbsp;-&nbsp;<a data-click="{'rsv_snapshot':'1'}" href="http://cache.baiducontent.com/c?m=9d78d513d9d430dc4f9b96690c66c0101843f4632bd6a0020edf843f96732b315011e0ac26520772d7d20d1016de4b4b9d862173471451c38cbe8c5dadbd855c5c9f2644676cf65661a70de88b182a9b66d618a8b81894a7ed77c4f28b94c854249a005e2cc7e78b2d51499572b44f66a6bbca18165413beea6336f908762bce2340b14cfbe3306e0581f4d85a4e937dd3361590ad62e16f13b604a46841621bfe&amp;p=9d3bc54ad5c44bf30eb8df2d021482&amp;newp=9f3bc54ad5c44bfd09be9b7c1c0c83231611d73f6590cf512496fe4893700d1a2a22b4fb66794d58dcc1766001ab4e5aeefa3475360425b791ca834fc9fdff6978ca28632c4ad410&amp;s=5ece19ab274c8fb8&amp;user=baidu&amp;fm=sc&amp;query=%CA%D0%D3%F2%C9%E7%BB%E1%D6%CE%C0%ED%CF%D6%B4%FA%BB%AF&amp;qid=aa09a0b600002ca9&amp;p1=10" target="_blank" class="m ">百度快 
照</a></div></div>
`

// var re_baidu_selenium = `<div[\d\D]*?class="result c-container "[\d\D]*?href="([^"]*?)"[\d\D]*?>([\d\D]*?)</a></h3>[\d\D]*?<div class="c-abstract"><span class=" newTimeFactor_before_abs  m">([^<]*?)&nbsp;-&nbsp;</span>([\d\D]*?)</div>`

// var re_baidu_selenium = `<div class="result c-container "[\d\D]*?href="([^"]*?)"[\d\D]*?>([\d\D]*?)</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">([^<]*?)&nbsp;-&nbsp;</span>([\d\D]*?)</div>`
func TestRe(t *testing.T) {
	re := regexp.MustCompile(re_baidu_selenium)
	match := re.FindAllSubmatch([]byte(body2), -1)

	for _, m := range match {
		// fmt.Println("match: ", string(m[0]))
		fmt.Println("url: ", string(m[1]))
		fmt.Println("title: ", string(m[2]))
		fmt.Println("updatetime: ", string(m[3]))
		fmt.Println("intro: ", string(m[4]))
	}

	// fmt.Print("match: ", match)
}
