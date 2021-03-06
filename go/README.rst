======================
วัดป่าโพธิญาณ (วัดเขื่อน)
======================

Create static website in Go_ programming language.

Development Environment:

  - `Ubuntu 17.04`_
  - `Go 1.8.1`_


Set Up Development Environment
++++++++++++++++++++++++++++++


1. `git clone`_ this repository:

   .. code-block:: bash

     # create a workspace in your home directory
     $ mkdir ~/dev
     # enter workspace
     $ cd ~/dev
     # git clone this repository
     $ git clone https://github.com/siongui/wat-pah-photiyan.git --depth=1
     # or clone with full depth
     #$ git clone https://github.com/siongui/wat-pah-photiyan.git


2. Update Ubuntu and install following packages:

   - Go_

   .. code-block:: bash

     $ cd ~/dev/wat-pah-photiyan
     $ make update_ubuntu
     $ make download_go
     $ make install


3. Run development server at http://localhost:8000/

   .. code-block:: bash

     $ make devserver


References
++++++++++

.. [1] `WAT NONG PAH PONG / วัดหนองป่าพง: Sakha 8 <https://watnsakha.blogspot.com/p/sakha-8.html>`_

.. [2] | `wat pah phothiyan - Google search <https://www.google.com/search?q=wat+pah+phothiyan>`_
       | `wat pah phothiyan - DuckDuckGo search <https://duckduckgo.com/?q=wat+pah+phothiyan>`_
       | `wat pah phothiyan - Ecosia search <https://www.ecosia.org/search?q=wat+pah+phothiyan>`_
       | `wat pah phothiyan - Bing search <https://www.bing.com/search?q=wat+pah+phothiyan>`_
       | `wat pah phothiyan - Yahoo search <https://search.yahoo.com/search?p=wat+pah+phothiyan>`_
       | `wat pah phothiyan - Baidu search <https://www.baidu.com/s?wd=wat+pah+phothiyan>`_
       | `wat pah phothiyan - Yandex search <https://www.yandex.com/search/?text=wat+pah+phothiyan>`_

.. [3] `พระโพธิญาณเถร (ชา สุภทฺโท) - วิกิพีเดีย <https://th.wikipedia.org/wiki/%E0%B8%9E%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93%E0%B9%80%E0%B8%96%E0%B8%A3_(%E0%B8%8A%E0%B8%B2_%E0%B8%AA%E0%B8%B8%E0%B8%A0%E0%B8%97%E0%B8%BA%E0%B9%82%E0%B8%97)>`_

.. [4] | `วัดป่าโพธิญาณ - Google search <https://www.google.com/search?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - DuckDuckGo search <https://duckduckgo.com/?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - Ecosia search <https://www.ecosia.org/search?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - Bing search <https://www.bing.com/search?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - Yahoo search <https://search.yahoo.com/search?p=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - Baidu search <https://www.baidu.com/s?wd=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_
       | `วัดป่าโพธิญาณ - Yandex search <https://www.yandex.com/search/?text=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

.. [5] | `Pure <http://purecss.io/>`_
       | `Bulma: a modern CSS framework based on Flexbox <http://bulma.io/>`_
       |   - `bulma themes - Google search <https://www.google.com/search?q=bulma+themes>`_
       |   - `Free Bulma templates <https://dansup.github.io/bulma-templates/>`_
       | `Basscss <http://basscss.com/>`_

.. [6] `GitHub - siongui/gotm: Go template manager <https://github.com/siongui/gotm>`_

.. [7] | `菩提智 阿姜查 - Google search <https://www.google.com/search?q=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - DuckDuckGo search <https://duckduckgo.com/?q=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - Ecosia search <https://www.ecosia.org/search?q=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - Bing search <https://www.bing.com/search?q=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - Yahoo search <https://search.yahoo.com/search?p=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - Baidu search <https://www.baidu.com/s?wd=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_
       | `菩提智 阿姜查 - Yandex search <https://www.yandex.com/search/?text=%E8%8F%A9%E6%8F%90%E6%99%BA+%E9%98%BF%E5%A7%9C%E6%9F%A5>`_

.. [8] `純CSS可切換選單 <https://siongui.github.io/zh/2017/02/21/css-only-toggle-menu/>`_

.. [9] | `golang yaml - Google search <https://www.google.com/search?q=golang+yaml>`_
       | `golang yaml - DuckDuckGo search <https://duckduckgo.com/?q=golang+yaml>`_
       | `golang yaml - Ecosia search <https://www.ecosia.org/search?q=golang+yaml>`_
       | `golang yaml - Qwant search <https://www.qwant.com/?q=golang+yaml>`_
       | `golang yaml - Bing search <https://www.bing.com/search?q=golang+yaml>`_
       | `golang yaml - Yahoo search <https://search.yahoo.com/search?p=golang+yaml>`_
       | `golang yaml - Baidu search <https://www.baidu.com/s?wd=golang+yaml>`_
       | `golang yaml - Yandex search <https://www.yandex.com/search/?text=golang+yaml>`_

.. [10] `GitHub - ghodss/yaml: A better way to marshal and unmarshal YAML in Golang <https://github.com/ghodss/yaml>`_

.. [11] | `งานมุฑิตาสักการะ - หลวงพ่อคำ นิสโสโก - วัดป่าไทยพัฒนา <https://www.facebook.com/Kottapan.Kaewsanga/posts/10211985490423478>`_
        | `เอนก ยสทินฺโน <https://www.facebook.com/Kottapan.Kaewsanga/posts/10212104349114871>`_

.. [12] | `นักรบ แห่งพุทธะ - พระอริยสงฆ์แห่งเมืองดอกบัวงาม... | Facebook <https://www.facebook.com/permalink.php?story_fbid=1961150254121265&id=100006789289467>`_
        | `นักรบ แห่งพุทธะ - หลวงพ่อบุญชู ฐิตคุโณ วัดป่าโพธิญาณ... | Facebook <https://www.facebook.com/photo.php?fbid=1961150117454612&set=a.1508900606012901.1073741829.100006789289467&type=3>`_

.. _Go: https://golang.org/
.. _Ubuntu 17.04: http://releases.ubuntu.com/17.04/
.. _Go 1.8.1: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _UNLICENSE: http://unlicense.org/
.. _go-libsass: https://github.com/wellington/go-libsass
.. _gettext-go: https://github.com/chai2010/gettext-go
