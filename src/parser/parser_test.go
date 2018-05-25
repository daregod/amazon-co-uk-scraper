package parser_test

import (
	"bytes"
	"testing"

	"github.com/daregod/amazon-co-uk-scraper/src/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Parser", func() {
	It("Compiles", func() {
		Expect(true).To(BeTrue())
	})
	It("Parse", func() {
		Expect(parser.Parse(bytes.NewBuffer(testData))).To(Equal(`£41.99`))
	})
})

func TestParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "parser")
}

var testData []byte = []byte(`
<div id="buybox" class="a-row a-spacing-medium">

<div id="rbbContainer" class="a-box-group dp-accordion"><div class="a-box a-first rbbSection selected dp-accordion-active"><div class="a-box-inner"><div class="a-section a-spacing-none a-padding-none"><div id="buyNewSection" class="a-section a-spacing-none a-padding-base rbbHeader dp-accordion-row">
<div class="a-row a-spacing-none a-grid-vertical-align a-grid-center"><div class="a-fixed-left-grid"><div class="a-fixed-left-grid-inner" style="padding-left:31px"><div class="a-text-left a-fixed-left-grid-col a-col-left" style="width:31px;margin-left:-31px;float:left">
<div class="a-row a-grid-vertical-align a-grid-center"><i class="a-icon a-icon-radio-active new-header"></i></div></div><div class="a-fixed-left-grid-col dualLineHeader a-col-right" style="padding-left:0%;float:left"><a class="a-link-normal rbbHeaderLink" href="javascript:void(0)">
<div class="a-section"><div class="a-row"><span class="a-text-bold">Buy New</span></div><div class="a-row"><span class="inlineBlock-display">

<span class="a-size-medium a-color-price offer-price a-text-normal">£41.99</span>

</span></div></div></a></div></div></div></div></div><div id="buyNewInner" class="rbbContent dp-accordion-inner"><div id="buyBoxInner" class="a-section a-spacing-small"><div class="a-row a-spacing-none"><div class="a-fixed-left-grid"><div class="a-fixed-left-grid-inner" style="padding-left:27px">
<div class="a-text-left a-fixed-left-grid-col a-col-left" style="width:27px;margin-left:-27px;float:left"></div><div class="a-fixed-left-grid-col a-col-right" style="padding-left:0%;float:left"><ul class="a-unordered-list a-nostyle a-vertical"><li><span class="a-list-item">
<div id="promiseBasedBadge_feature_div" data-feature-name="promiseBasedBadge" class="a-section a-spacing-none feature"></div></span></li></ul></div></div></div></div></div><div class="a-section a-spacing-small a-spacing-top-micro"><div class="a-row"><span class="a-color-base buyboxShippingLabel">
<a href="/gp/help/customer/display.html/ref=mk_sss_dp_1?ie=UTF8&amp;pop-up=1&amp;nodeId=202094700" target="AmazonHelp" onclick="return amz_js_PopWin(this.href,'AmazonHelp','width=550,height=550,resizable=1,scrollbars=1,toolbar=0,status=0')">FREE Delivery</a> in the UK.</span></div></div><div class="a-section a-spacing-none"></div><div class="a-section a-spacing-small">

<div id="availability" class="a-section a-spacing-none">
<span class="a-size-medium a-color-success">
In stock.
</span>
</div>

<div class="a-section a-spacing-none"></div><div id="merchant-info" class="a-section a-spacing-mini">Dispatched from and sold by Amazon.<span class="">Gift-wrap available.</span></div></div><div class="a-section a-spacing-small"><div id="selectQuantity" class="a-section a-spacing-none a-padding-none"><span class="a-declarative" data-action="quantity-dropdown" data-quantity-dropdown="{}">
<div class="a-row a-spacing-base"><div class="a-column a-span12 a-text-left"><span class="a-dropdown-container"><label for="quantity" class="a-native-dropdown">Quantity:</label><select name="quantity" autocomplete="off" id="quantity" tabindex="-1" class="a-native-dropdown">
<option value="1" selected="">1</option><option value="2">2</option><option value="3">3</option><option value="4">4</option><option value="5">5</option><option value="6">6</option><option value="7">7</option><option value="8">8</option>
<option value="9">9</option><option value="10">10</option><option value="11">11</option><option value="12">12</option><option value="13">13</option><option value="14">14</option><option value="15">15</option><option value="16">16</option>
<option value="17">17</option><option value="18">18</option><option value="19">19</option><option value="20">20</option><option value="21">21</option><option value="22">22</option><option value="23">23</option><option value="24">24</option>
<option value="25">25</option><option value="26">26</option><option value="27">27</option><option value="28">28</option><option value="29">29</option><option value="30">30</option></select>
<span tabindex="-1" class="a-button a-button-dropdown a-button-small" id="a-autoid-0" style="min-width:0"><span class="a-button-inner"><span class="a-button-text a-declarative" data-action="a-dropdown-button" role="button" tabindex="0" aria-hidden="true" id="a-autoid-0-announce"><span class="a-dropdown-label">Quantity:</span>
<span class="a-dropdown-prompt">1</span></span><i class="a-icon a-icon-dropdown"></i></span></span></span></div></div></span></div></div><div id="bbopAndCartBox" class="a-box"><div class="a-box-inner">
<div class="a-button-stack"><span id="submit.add-to-cart" class="a-button a-spacing-small a-button-primary a-button-icon"><span class="a-button-inner"><i class="a-icon a-icon-cart"></i><input id="add-to-cart-button" name="submit.add-to-cart" title="Add to Shopping Basket" data-hover="Select <b>__dims__</b> from the left<br> to add to Basket" class="a-button-input" value="Add to Basket" aria-labelledby="submit.add-to-cart-announce" type="submit">
<span id="submit.add-to-cart-announce" class="a-button-text" aria-hidden="true">Add to Basket</span></span></span></div></div></div><div class="a-row a-spacing-none"><div id="oneClickSignIn" class="a-section a-spacing-none"><div class="a-divider a-divider-break a-spacing-micro"><h5><a href="/gp/product/utility/edit-one-click-pref.html?ie=UTF8&amp;query=selectObb%3dnew&amp;returnPath=%2fgp%2fproduct%2f1787125645"><span class="a-size-mini">Turn on 1-Click ordering for this browser</span></a></h5></div></div></div>
<div class="a-row"><div id="dpFastTrack_feature_div" data-feature-name="dpFastTrack" data-template-name="dpFastTrack" class="a-section a-spacing-none a-spacing-top-small feature"></div></div><hr class="a-divider-normal"><span class="a-declarative" data-action="dpContextualIngressPt" data-dpcontextualingresspt="{}"><a class="a-link-normal" href="#"><div class="a-row a-spacing-mini"><div class="a-column a-span12 a-text-left"><div id="contextualIngressPt">
<div id="contextualIngressPtPin"></div><span id="contextualIngressPtLabel" class="a-size-small"><div id="contextualIngressPtLabel_deliveryShortLine">Select delivery location</div></span></div></div></div></a></span><div class="a-row"><div id="holidayAvailabilityMessage_feature_div" data-feature-name="holidayAvailabilityMessage" data-template-name="holidayAvailabilityMessage" class="a-section a-spacing-top-small feature"></div></div></div></div></div></div>
<div class="a-box a-last rbbSection unselected"><div class="a-box-inner"><div class="a-section a-spacing-none a-padding-none"><div id="usedBuySection" class="a-section a-spacing-none a-padding-base rbbHeader dp-accordion-row"><div class="a-row a-spacing-none a-grid-vertical-align a-grid-center"><div class="a-fixed-left-grid"><div class="a-fixed-left-grid-inner" style="padding-left:31px"><div class="a-text-left a-fixed-left-grid-col a-col-left" style="width:31px;margin-left:-31px;float:left">
<div class="a-row a-grid-vertical-align a-grid-center"><i class="a-icon a-icon-radio-inactive new-header"></i></div></div><div class="a-fixed-left-grid-col dualLineHeader a-col-right" style="padding-left:0%;float:left"><a class="a-link-normal rbbHeaderLink" href="javascript:void(0)">
<div class="a-section"><div class="a-row"><span class="a-text-bold">Buy Used</span></div><div class="a-row"><span class="inlineBlock-display">

<span class="a-color-base offer-price a-text-normal">£37.49</span>

</span></div></div></a></div></div></div></div></div><div id="usedbuyBox" class="rbbContent dp-accordion-inner" style="display:none">
<input id="usedMerchantID" name="usedMerchantID" value="AEUF4A7MV0O5S" type="hidden"><input id="usedOfferListingID" name="usedOfferListingID" value="nYOm5b3TuCU0I2Co%2FpNKWHAppIBI1bSy00crMIIRNGKic2N6f9klVpORR4L1ZgRrrXxNyOSFlrVJRWTnXIdUsrTQmCGihHJlwPbeD4FonYV%2BFysj0QcltrYm%2BL2jGoQcgCM2U%2BRoL8AllSKyS%2BZiS2p286gODSLa" type="hidden">
<input id="usedSellingCustomerID" name="usedSellingCustomerID" value="AEUF4A7MV0O5S" type="hidden"><div class="a-section a-spacing-mini"><div class="a-row">+&nbsp;£0.00&nbsp;delivery</div></div><div class="a-section a-spacing-base"><div class="a-row"><strong>Used: Like New</strong><span class="a-size-base"><span class="a-color-tertiary"> | </span><a id="usedItemConditionInfoLink" class="a-link-normal a-declarative" href="#">Details</a></span></div><div class="a-row">Sold by<a class="a-link-normal" href="/gp/help/seller/at-a-glance.html?ie=UTF8&amp;seller=AEUF4A7MV0O5S">IndiGlobalShop</a></div></div>
<div class="a-popover-preload" id="a-popover-usedItemConditionDetailsPopover"><div class="a-section a-spacing-micro"><span class="a-size-mini"><strong>Condition:</strong>Used: Like New</span></div></div><div class="accessCode-spacing"></div><span class="a-declarative" data-action="dpContextualIngressPt" data-dpcontextualingresspt="{}"><a class="a-link-normal" href="#"><div class="a-row a-spacing-mini"><div class="a-column a-span12 a-text-left"><div id="contextualIngressPt"><div id="contextualIngressPtPin"></div><span id="contextualIngressPtLabel" class="a-size-small"><div id="contextualIngressPtLabel_deliveryShortLine">Select delivery location</div></span></div></div></div></a></span>
<div class="a-button-stack"><span id="submit.add-to-cart-ubb" class="a-button a-spacing-small a-button-primary a-button-icon"><span class="a-button-inner"><i class="a-icon a-icon-cart"></i><input id="add-to-cart-button-ubb" name="submit.add-to-cart-ubb" title="Add to Shopping Basket" data-hover="Select <b>__dims__</b> from the left<br> to add to Basket" class="a-button-input" value="Add to Basket" aria-labelledby="submit.add-to-cart-ubb-announce" type="submit"><span id="submit.add-to-cart-ubb-announce" class="a-button-text" aria-hidden="true">Add to Basket</span></span></span></div>
<div class="a-section a-spacing-none a-text-center"><div class="a-row a-spacing-none"><div id="oneClickSignInUBB" class="a-section a-spacing-none"><div class="a-divider a-divider-break a-spacing-micro"><h5><a href="/gp/product/utility/edit-one-click-pref.html?ie=UTF8&amp;query=selectObb%3dused&amp;returnPath=%2fgp%2fproduct%2f1787125645"><span class="a-size-mini">Turn on 1-Click ordering for this browser</span></a></h5></div></div></div></div></div></div></div></div></div></div>
`)
