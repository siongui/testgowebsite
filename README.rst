======================
วัดป่าโพธิญาณ (วัดเขื่อน)
======================

Create static website in Go_ programming language.

Development Environment:

  - `Ubuntu 16.10`_
  - `Go 1.7.5`_


Set Up Development Environment
++++++++++++++++++++++++++++++


1. `git clone`_ this repository:

   .. code-block:: bash

     # create a workspace in your home directory
     $ mkdir ~/dev
     # enter workspace
     $ cd ~/dev
     # git clone this repository
     $ git clone https://github.com/siongui/testgowebsite.git --depth=1
     # or clone with full depth
     #$ git clone https://github.com/siongui/testgowebsite.git


2. Update Ubuntu and install following packages:

   - Go_

   .. code-block:: bash

     $ cd ~/dev/testgowebsite
     $ make update_ubuntu
     $ make download_go


3. Run development server at http://localhost:8000/

   .. code-block:: bash

     $ make devserver


References
++++++++++

.. [1] `WAT NONG PAH PONG / วัดหนองป่าพง: Sakha 8 <https://watnsakha.blogspot.com/p/sakha-8.html>`_

.. [2] `wat pah phothiyan - Google search <https://www.google.com/search?q=wat+pah+phothiyan>`_

       `wat pah phothiyan - DuckDuckGo search <https://duckduckgo.com/?q=wat+pah+phothiyan>`_

       `wat pah phothiyan - Bing search <https://www.bing.com/search?q=wat+pah+phothiyan>`_

       `wat pah phothiyan - Yahoo search <https://search.yahoo.com/search?p=wat+pah+phothiyan>`_

       `wat pah phothiyan - Baidu search <https://www.baidu.com/s?wd=wat+pah+phothiyan>`_

       `wat pah phothiyan - Yandex search <https://www.yandex.com/search/?text=wat+pah+phothiyan>`_

.. [3] `พระโพธิญาณเถร (ชา สุภทฺโท) - วิกิพีเดีย <https://th.wikipedia.org/wiki/%E0%B8%9E%E0%B8%A3%E0%B8%B0%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93%E0%B9%80%E0%B8%96%E0%B8%A3_(%E0%B8%8A%E0%B8%B2_%E0%B8%AA%E0%B8%B8%E0%B8%A0%E0%B8%97%E0%B8%BA%E0%B9%82%E0%B8%97)>`_

.. [4] `วัดป่าโพธิญาณ - Google search <https://www.google.com/search?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

       `วัดป่าโพธิญาณ - DuckDuckGo search <https://duckduckgo.com/?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

       `วัดป่าโพธิญาณ - Bing search <https://www.bing.com/search?q=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

       `วัดป่าโพธิญาณ - Yahoo search <https://search.yahoo.com/search?p=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

       `วัดป่าโพธิญาณ - Baidu search <https://www.baidu.com/s?wd=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

       `วัดป่าโพธิญาณ - Yandex search <https://www.yandex.com/search/?text=%E0%B8%A7%E0%B8%B1%E0%B8%94%E0%B8%9B%E0%B9%88%E0%B8%B2%E0%B9%82%E0%B8%9E%E0%B8%98%E0%B8%B4%E0%B8%8D%E0%B8%B2%E0%B8%93>`_

.. [5] | `Pure <http://purecss.io/>`_
       | `Bulma: a modern CSS framework based on Flexbox <http://bulma.io/>`_
       | `Basscss <http://basscss.com/>`_


.. _Go: https://golang.org/
.. _Ubuntu 16.10: http://releases.ubuntu.com/16.10/
.. _Go 1.7.5: https://golang.org/dl/
.. _git clone: https://www.google.com/search?q=git+clone
.. _UNLICENSE: http://unlicense.org/
.. _go-libsass: https://github.com/wellington/go-libsass
.. _gettext-go: https://github.com/chai2010/gettext-go
