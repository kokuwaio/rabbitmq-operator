FROM gitpod/workspace-full

USER root

### Docker client ###
RUN curl -fsSL https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_Unstable/Release.key | apt-key add - \
    && curl -fsSL https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_Testing/Release.key | apt-key add - \
    && curl -fsSL https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_10/Release.key | apt-key add - \
    && echo 'deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_Unstable/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list \
    && echo 'deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_Testing/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list \
    && echo 'deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/Debian_10/ /' > /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list \
    && apt-get update -qq \
    && apt-get -qq -y install podman \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*

ADD https://raw.githubusercontent.com/containers/buildah/master/contrib/buildahimage/stable/containers.conf /etc/containers/

# Adjust storage.conf to enable Fuse storage.
RUN chmod 644 /etc/containers/containers.conf; sed -i -e 's|^#mount_program|mount_program|g' -e '/additionalimage.*/a "/var/lib/shared",' -e 's|^mountopt[[:space:]]*=.*$|mountopt = "nodev,fsync=0"|g' /etc/containers/storage.conf
RUN mkdir -p /var/lib/shared/overlay-images /var/lib/shared/overlay-layers /var/lib/shared/vfs-images /var/lib/shared/vfs-layers; touch /var/lib/shared/overlay-images/images.lock; touch /var/lib/shared/overlay-layers/layers.lock; touch /var/lib/shared/vfs-images/images.lock; touch /var/lib/shared/vfs-layers/layers.lock

USER gitpod
# Install custom tools, runtime, etc.
