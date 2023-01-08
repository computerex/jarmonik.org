+++
title = Home
render = true
+++

<div class='recent-posts'>
    {{home_recent_posts}}
</div>

<h2><br>LunarTransferMFD 1.6</h2>
<h2><small>Updated 24.12.2019</small> </h2>
This is MFD addon for <a href="https://web.archive.org/web/20220521025709/http://orbit.medphys.ucl.ac.uk/"> Orbiter space flight simulator</a> version 2010, 2016
<p>
LunarTransferMFD is aa numerical lunar transfer trajectory calculator. It can be used
to compute single impulse lunar transfer trajectories with better accuracy than
typical patched conic applications such as IMFD. LunarTransferMFD is based on Broyden's
method with numerical forward trajectory model (RKF56). Typical user input parameters are
desired time of periapis passage, altitude and flight heading in the periapis.

</p><p>In the current version, operation is limited in Apollo style flights 
from the Earth's surface to lunar surface and back. Missions starting from unaligned orbits are not yet
supported and many other advanced features are still under work.

</p><p>LunarTransferMFD can be used only with high thrust vehicles such as Apollo and LRO. Low thrust
propulsion systems like ion-engines won't work with LTMFD. 
<br><br><br>
<a href="https://orbiter-mods.com/mod/47">Download 1.6</a> (~788k) For Orbiter 2016.<br>
<a href="https://orbiter-mods.com/mod/46">Download 1.4</a> (~788k) For Orbiter 2010.<br>
<br>
LunarTransferMFD is using SSE2 instruction set by default. Some older CPUs like Athlon XP doesn't support these instructions.

<h2><br>IMFD - Space Navigation Computer Version 5.7</h2>
<h2><small>Updated 17.12.2016</small></h2>

This is MFD addon for <a href="https://web.archive.org/web/20220521025709/http://orbit.medphys.ucl.ac.uk/"> Orbiter space flight simulator</a> versions 2006 - 2016.

<b>Key Features of the IMFD:</b><br>
<ul>
<li>Map program can be used to display entire solar system and predict the actual course of the vessel with numerical trajectory calculator. The predictor can be also used to predict a hypothetical scenarios as well. This is very useful when planning a free-return trajectories.<br><br>

</li><li>Course program is a collection of minor navigation programs those will let you to navigate through the solar system 
to other planets, moons and comets. Primary programs are based on the lambert solver. In the most simplest case only a date
of ejection and the Date of arrival is required as an user input.<br><br>


</li><li>BaseApproach program is designed for Apollo style Moon to Earth return missions. It can be used to target
a bases on the Mars and land directly without an orbit insertion. The program will also allow to establish an orbit right over a specific base on a planet so that it is possible to continue the descent on the base.
That's very useful when landing on the Moon.<br><br>

</li><li>Surface Launch program is usefull when starting interplanetary or lunar missions from the planet surface. The program will compute the time to launch window and required launch heading. This program is the only one that doesn't compute any burn guidance data. A vessel specific ascent profile should be used to reach low orbit.<br><br>

</li><li>Orbit Eject program is required when starting interplanetary missions from the orbit. This program is used in a combination with the other programs to convert the escape information into burn guidance data. 
</li></ul>
<br>

<br>
Manual of the 4.2.2 release may contain some helpful information that is not included in later manuals. Installation of 4.2.2 doesn't remove existing installations of new versions.<br> 
<br> 
<a href="https://orbiter-mods.com/mod/48">Download 5.7</a> (~1M) (for Orbiter 2016)<br>

<h2>D3D9Client</h2>
Release Date 5-Aug-2021
<hr>
<h2>Installation</h2>
This is a graphics client/engine project to <a href="http://orbit.medphys.ucl.ac.uk/">Orbiter Spaceflight Simulator</a>. To install the client you need to download a package from a list below and extract it in the <b>root</b> folder of the Orbiter.
Any previously existing files should be replaced. Also to use a graphics client in Orbiter you need to run "Orbiter_ng.exe" instead 
of "Orbiter.exe" which uses a build in DX7 Engine. The client must be activated from the Modules tab otherwise you will see
a command console while trying to start a scenario. Read further information from /Doc/D3D9Client.pdf
<br><br>D3D9Client development and support thread is located in <a href="https://www.orbiter-forum.com/index.php">Orbiter-Forum</a>
<br><br><hr>
<h2>DirectX Runtimes</h2>
You will need a DirectX June 2010 redistributable backage to run the client. If the redistributable package 
isn't installed in your computer you will receive an error message "The program can't start because d3dx9_43.dll is
missing from your computer". Or you may see a pop-up window in Orbiter LaunchPad telling about a missing runtimes. 
If that happens then download the package from a link below and extract the content of the package in any empty directory 
you want and then find a <b>Setup.exe</b> and <b>run it.</b><i> It is a common mistake to forget the run the Setup.exe</i>
You can delete the contents of the directory after the setup is completed. The directory is just a
temporary storage for the installation files.<br><br>
Here is a link: <a href="https://www.microsoft.com/en-us/download/details.aspx?id=8109">June 2010 Redistributables</a>  
<br><br><hr>
<h2>Terms of Use and Warranty</h2>
You can use the client and other downloads found from this site under the same terms as the Orbiter itself. <a href="http://orbit.medphys.ucl.ac.uk/terms.html">Terms of Use</a><br><br>
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR
IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
<br><br><hr>
<h2>Orbiter Beta</h2>
When using a graphics client with Orbiter Beta, a special attention should be paid that orbiter and client versions do match properly,
otherwise the client might not run. The Orbiter version is denoted in a file name, as an example in a name "D3D9ClientBeta28.5-forBETA r84(r1054)"
targeted orbiter version (or revision) is "r84".

<h2>D3D9Client Downloads</h2>
<a href="https://orbiter-mods.com/mod/12">Download for Orbiter 2016</a><br>
<a href="https://orbiter-mods.com/mod/45">Download for Orbiter 2010</a><br>

<h3>MicroTexture Pack</h3>
This is a planetary surface microtexture package for D3D9Client.
The package contains textures for the Moon and the Mars. 
Here is a screen shot before and after installation.
<br><br>
<a href="https://orbiter-mods.com/mod/44">Download (~25MB)</a><br><br>   
