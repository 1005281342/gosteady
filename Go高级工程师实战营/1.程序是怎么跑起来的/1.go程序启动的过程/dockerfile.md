FROM centos

RUN cd /etc/yum.repos.d/ \

&& sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* \

&& sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* \

&& yum update -y \

&& yum install golang -y \ 

&& yum install dlv -y \

&& yum install binutils -y \ 

&& yum install vim -y \ 

&& yum install gdb -y

