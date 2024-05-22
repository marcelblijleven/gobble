FROM python:3.12.2 as base

WORKDIR /app

ENV PYTHONUNBUFFERED=1
ENV POETRY_VERSION=1.8.2
ENV POETRY_HOME=/opt/poetry
ENV POETRY_NO_INTERACTION=1
ENV POETRY_VIRTUALENVS_CREATE=false
ENV POETRY_VENV=/opt/poetry-venv

ENV PATH="$POETRY_HOME/bin:$POETRY_VENV/bin:$PATH"

# Create virtualenv for poetry
RUN python3 -m venv $POETRY_VENV
# Set Python path
ENV PYTHONPATH="/app:$PYTHONPATH"

# Install poetry into the virtualenv
RUN $POETRY_VENV/bin/pip install -U pip setuptools
RUN $POETRY_VENV/bin/pip install -U poetry==${POETRY_VERSION}

FROM base AS runner

WORKDIR /app

COPY README.md poetry.lock pyproject.toml ./
COPY gobble gobble

# Install runtime dependencies with Poetry
RUN poetry config --list
RUN poetry install --without dev

# Set user group
RUN chgrp -R 0 /app && \
    chmod -R g=u /app && \
    chmod -R g+w /app
# Run as user
USER $USER

# Python specific variables
# Do not buffer output
ENV PYTHONUNBUFFERED=1
CMD ["poetry", "run", "uvicorn", "gobble.app:app", "--host", "0.0.0.0"]