package engine

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
	// t.Skip()
	rePublisher := regexp.MustCompile(price1)
	// mPublisher := rePublisher.FindSubmatch([]byte(all3))
	mPublisher := rePublisher.FindSubmatch([]byte(all3))

	fmt.Println("match: ", string(mPublisher[1]))
}
