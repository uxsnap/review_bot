FROM python:3.9

SHELL ["/bin/bash", "-c"]

ARG PORT
ARG DP_CONFIG_NAME

RUN python3 -m venv env
RUN source env/bin/activate
RUN pip install deeppavlov
RUN python -m deeppavlov install ${DP_CONFIG_NAME}  
RUN python -m deeppavlov download ${DP_CONFIG_NAME}

CMD source env/bin/activate
CMD python -m deeppavlov riseapi ${DP_CONFIG_NAME} -p ${PORT}