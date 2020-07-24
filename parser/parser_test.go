package parser

import (
	"fmt"
	"regexp"
	"testing"
)

const (
	url     = `https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C`
	pattern = `<a href="([^"]+)" title="([^"]+)" [\d\D]+?onclick="moreurl`
	str     = `<a href="https://book.douban.com/subject/30192800/" title="Python神经网络编程" onclick="moreurl(this,{i:'2',query:'',subject_id:'30192800',from:'book_subject_search'})">`

	all = `<div id="info" class="">



    
    
  
    <span>
      <span class="pl"> 作者</span>:
        
            
            <a class="" href="/search/%E9%82%B1%E9%94%A1%E9%B9%8F">邱锡鹏</a>
    </span><br>

    
    
  
    <span class="pl">出版社:</span> 机械工业出版社<br>

    
    
  
    <span class="pl">出品方:</span>&nbsp;<a href="https://book.douban.com/series/48838?brand=1">华章IT</a><br>

    
    
  

    
    
  

    
    
  

    
    
  
    <span class="pl">出版年:</span> 2020-4-10<br>

    
    
  
    <span class="pl">页数:</span> 448<br>

    
    
  
    <span class="pl">定价:</span> 149.00元<br>

    
    
  
    <span class="pl">装帧:</span> 平装<br>

    
    
  
    <span class="pl">丛书:</span>&nbsp;<a href="https://book.douban.com/series/51864">人工智能技术丛书</a><br>

    
    
  
    
      
      <span class="pl">ISBN:</span> 9787111649687<br>


</div>`

	all2 = `<div id="interest_sectl" class="">
  <div class="rating_wrap clearbox" rel="v:rating">
    <div class="rating_logo">豆瓣评分</div>
    <div class="rating_self clearfix" typeof="v:Rating">
      <strong class="ll rating_num " property="v:average"> 9.6 </strong>
      <span property="v:best" content="10.0"></span>
      <div class="rating_right ">
          <div class="ll bigstar50"></div>
            <div class="rating_sum">
                <span class="">
                    <a href="collections" class="rating_people"><span property="v:votes">621</span>人评价</a>
                </span>
            </div>


      </div>
    </div>
          
            
            
<span class="stars5 starstop" title="力荐">
    5星
</span>

            
<div class="power" style="width:64px"></div>

            <span class="rating_per">77.9%</span>
            <br>
            
            
<span class="stars4 starstop" title="推荐">
    4星
</span>

            
<div class="power" style="width:15px"></div>

            <span class="rating_per">18.4%</span>
            <br>
            
            
<span class="stars3 starstop" title="还行">
    3星
</span>

            
<div class="power" style="width:2px"></div>

            <span class="rating_per">3.2%</span>
            <br>
            
            
<span class="stars2 starstop" title="较差">
    2星
</span>

            
<div class="power" style="width:0px"></div>

            <span class="rating_per">0.2%</span>
            <br>
            
            
<span class="stars1 starstop" title="很差">
    1星
</span>

            
<div class="power" style="width:0px"></div>

            <span class="rating_per">0.3%</span>
            <br>
    </div>
</div>`

	all3 = `<div id="info" class="">
    <span>
      <span class="pl"> 作者</span>:
            <a class="" href="/search/%E9%82%B1%E9%94%A1%E9%B9%8F">邱锡鹏</a>
    </span><br/>
    <span class="pl">出版社:</span> 机械工业出版社<br/>
    <span class="pl">出版年:</span> 2020-4-10<br/>
    <span class="pl">页数:</span> 448<br/>
    <span class="pl">定价:</span> 149.00元<br/>`

	name1 = `<h1>[\d\D]*?<span.*?>([^<]+)</span>[\d\D]*?</h1>`

	author1 = `<span.*?> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`

	publisher1 = `<span class="pl">出版社:</span> ([^<]+)<br>`

	publishTime1 = `<span class="pl">出版年:</span> ([^<]+)<br>`

	price1 = `<span.*?>定价:</span> ([^<]+)<br>`

	score1 = `<strong class="ll rating_num ".*?>([^<]+) </strong>`
)

func TestParser(t *testing.T) {
	t.Skip()
	rePublisher := regexp.MustCompile(price1)
	// mPublisher := rePublisher.FindSubmatch([]byte(all3))
	mPublisher := rePublisher.FindSubmatch([]byte(all3))

	fmt.Println("match: ", string(mPublisher[1]))
}

var pageStr = `
<li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/26582822/" onclick="moreurl(this,{i:'19',query:'',subject_id:'26582822',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s28278604.jpg" width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/26582822/" title="重新定义公司" onclick="moreurl(this,{i:'19',query:'',subject_id:'26582822',from:'book_subject_search'})">

    重新定义公司


    
      <span style="font-size:12px;"> : 谷歌是如何运营的 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美]埃里克·施密特 / 靳婷婷、陈序、何晔 / 中信出版社 / 2015-8 / 49.00

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar40"></span>
        <span class="rating_nums">8.2</span>

    <span class="pl">
        (6597人评价)
    </span>
  </div>



        
  
  
  
    <p>谷歌高管手绘风漫画视频：
http://v.youku.com/v_show/id_XMTMxMzQ3NjMyMA==.html?from=y1.7-1.2... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/26582822/buylinks">
        纸质版 36.80 元起
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/46366039/?dcs=tag-buylink&amp;dcm=douban&amp;dct=26582822">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>

<div class="paginator">
        <span class="prev">
            &lt;前页
        </span>
        
        

                <span class="thispage">1</span>
                
            <a href="/tag/股票?start=20&amp;type=T">2</a>
        
                
            <a href="/tag/股票?start=40&amp;type=T">3</a>
        
                
            <a href="/tag/股票?start=60&amp;type=T">4</a>
        
                
`

func TestPage(t *testing.T) {
	t.Skip()
	// reLevel2Page := `<span class="thispage">1</span>[\d\D]*?<a href="([^"]+)">\d</a>[\d\D]*`
	reLevel2Flag := `<span class="thispage">2</span>`
	re := regexp.MustCompile(reLevel2Flag)
	// mPublisher := rePublisher.FindSubmatch([]byte(all3))
	urlList := re.FindAllSubmatch([]byte(pageStr), -1)

	fmt.Println("urlList == nil: ", urlList == nil)

	for _, url := range urlList {
		fmt.Println("match: ", string(url[0]))
	}

}

var url1 = "https://book.douban.com/subject/30293801/"

func TestUrlId(t *testing.T) {
	t.Skip()
	reUrlId := `([\d]+)`
	re := regexp.MustCompile(reUrlId)
	// mPublisher := rePublisher.FindSubmatch([]byte(all3))
	id := re.FindString(url1)

	fmt.Println("id: ", id)
}

var cate1 = `<div id="content">
    
    <h1>豆瓣音乐标签: OST</h1>

    <div class="grid-16-8 clearfix">
        
        
        <div class="article">`

func TestCate(t *testing.T) {
	reCate := `<h1>.*?: ([^<]+)</h1>`
	re := regexp.MustCompile(reCate)
	cate := re.FindSubmatch([]byte(cate1))

	fmt.Println("cate: ", string(cate[1]))
}
