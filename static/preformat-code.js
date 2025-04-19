// Source: https://gist.github.com/chriscauley/7b2c1227aae330d3f0dc5735a697c008

/*
Removes the minimum leading whitespace for each line in a pre > code tag.
Also optionally escapes html if you include the "nuke-html" class
When writing HTML snippets for slides or blogs, it's a pain to have to do this:

	</div> <!-- original indentation level -->
<pre><code class="html">&lt;ul class="demo"&gt;
  &lt;li&gt;No order here&lt;/li&gt;
  &lt;li&gt;Or here&lt;/li&gt;
  &lt;li&gt;Or here&lt;/li&gt;
  &lt;li&gt;Or here&lt;/li&gt;
&lt;/ul&gt;</pre></code>

What you really want is this:

	</div>
	<pre><code class="html nuke-html">
	    <ul class="demo">
	      <li>No order here</li>
	      <li>Or here</li>
	      <li>Or here</li>
	      <li>Or here</li>
	    </ul>
	</code></pre>
*/

document.querySelectorAll("pre > code").forEach(function(element, n) {
	if (element.classList.contains("nuke-html")) {
		var text = element.innerHTML;
	} else {
		var text = element.innerText;
	}
	text = text.replace(/^\n/,'').trimEnd(); // goodbye starting whitespace

	// Normalise indentation
	var to_kill = Infinity;
	var lines = text.split("\n");
	// Finds the minimum indentation level
	for (var i = 0; i < lines.length; i++) {
		if (!lines[i].trim()) { continue; }
		to_kill = Math.min(lines[i].search(/\S/), to_kill);
	}
	out = [];
	// Removes the minimum indentation level
	for (var i = 0; i < lines.length; i++) {
		out.push(lines[i].replace(new RegExp("^\\s{" + to_kill + "}", "g"), ""));
	}
	element.innerHTML = out.join("\n");
});
