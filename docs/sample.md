
This is a sample markdown document.

* a list
* of things
* is here
* with a link to www.google.com.

> This is a quote.

1. a numbered
2. list
3. is fancy


```
func (mds *MarkdownServer) tryFileAndIndexMD(urlPath string) (http.File, bool) {

	f, canRead := mds.openWithMDSuffix(urlPath)
	if canRead {
		return f, canRead
	}

	return mds.openWithMDSuffix(urlPath + "/index")
}
```

<strong>this is bold</strong>

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur enim urna,
feugiat et finibus eget, tincidunt ut tellus. Vivamus euismod facilisis mi, et
euismod ligula vehicula in. Aliquam erat volutpat. Etiam eget condimentum
dolor. Integer non ornare purus. Cras efficitur quam id metus ullamcorper
malesuada. Donec vehicula nunc ac elit interdum, vitae vulputate neque
eleifend. Cras ut sapien non mi pellentesque interdum.

Nunc pretium tristique ornare. Suspendisse non euismod magna. Sed eget augue in
quam accumsan ultricies ut euismod sapien. Cras mauris dui, consequat sed
sodales vel, gravida in mauris. Donec eget interdum massa, et rutrum odio.
Donec vestibulum ex non massa euismod, at mollis mauris vestibulum. Fusce risus
arcu, finibus sed sagittis non, pulvinar et neque. Maecenas eleifend, ex sit
amet interdum pretium, lectus libero efficitur ipsum, sed maximus justo quam
nec sem. Mauris eget sapien varius, sodales elit vel, vulputate velit. Donec
vehicula nisi at diam sollicitudin interdum. Aliquam sit amet fringilla mauris,
ut rhoncus magna. Pellentesque interdum malesuada augue. Donec augue lectus,
ullamcorper vestibulum nunc at, ultrices ultrices erat.

Phasellus vel posuere lorem, quis venenatis ipsum. Nam ut sodales ante. Aliquam
faucibus eros lectus, ut tincidunt orci ullamcorper a. Sed ultrices, sem eu
eleifend placerat, mi massa dapibus ligula, ac condimentum eros eros id magna.
Duis tempus, felis sed aliquam eleifend, mi magna ultrices diam, a mollis elit
metus in odio. Vivamus ullamcorper dapibus fringilla. Pellentesque pharetra
tincidunt urna a fermentum. Nunc vulputate nibh at purus congue, id condimentum
metus tristique. Duis magna neque, vestibulum sit amet pulvinar vitae, pharetra
non mi. Pellentesque habitant morbi tristique senectus et netus et malesuada
fames ac turpis egestas. Ut sed purus ac massa facilisis maximus gravida eu
quam. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere
cubilia curae;

Mauris pretium erat sapien, non tincidunt eros consectetur et. Class aptent
taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.
Maecenas tortor urna, luctus et velit in, vestibulum tempor ex. Nulla ut lacus
eget mi aliquam cursus et vitae elit. Phasellus non arcu sed mi pulvinar
commodo eu at lorem. Donec quis rutrum lorem. Proin viverra, lectus nec posuere
cursus, odio augue aliquam turpis, ut suscipit velit orci a tellus. Vestibulum
ullamcorper, orci vitae tristique ultrices, libero tellus aliquet lacus, id
ultricies odio magna nec tellus. Ut ut nulla sit amet mauris tincidunt
scelerisque. Ut pharetra interdum leo, quis auctor est. Sed fringilla libero id
vestibulum pulvinar. Orci varius natoque penatibus et magnis dis parturient
montes, nascetur ridiculus mus. Nam non libero ac dui consequat dapibus. Proin
sed ante urna.

Quisque condimentum vulputate justo in iaculis. Praesent nunc nulla, malesuada
sed ipsum a, semper dignissim elit. Morbi metus lacus, dictum quis gravida et,
cursus eget lectus. Curabitur venenatis feugiat dictum. Sed nec orci lorem.
Fusce congue tincidunt neque, id suscipit magna convallis in. Vestibulum congue
nec leo id euismod. Cras eu facilisis magna. Maecenas egestas metus nibh, sed
vulputate velit pulvinar non. Donec eu venenatis felis, ut volutpat tortor.
Aliquam efficitur dictum est in tempus. Maecenas cursus nisl eget malesuada
molestie. Aenean pulvinar nisl vitae iaculis congue. Nunc sed dignissim lectus,
in feugiat est.

